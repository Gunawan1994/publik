package cache

import (
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

// NetworkTCP for tcp
const NetworkTCP = "tcp"

// Redis defines a redis agent
type Redis struct {
	pool *redigo.Pool
}

// RedisOptions are options for the redis agent
type RedisOptions struct {
	MaxIdle   int
	MaxActive int
	Timeout   int
	Wait      bool
}

// RedisConfig config for redis agent
type RedisConfig struct {
	Address string
	AppName string
	Network string
	Options []RedisOptions
}

// dialFunc is func used to connect to the redis server
type dialFunc func(network, addr string, opts ...redigo.DialOption) (redigo.Conn, error)

// NewRedis creates new redis agent object
func NewRedis(cfg RedisConfig) *Redis {
	return newRedisWithDialer(cfg, redigo.Dial)
}

//func newAgentWithDialer(cfg Config, dialer dialFunc, opts ...RedisOptions) (redis *Redis) {
func newRedisWithDialer(cfg RedisConfig, dialer dialFunc) (redis *Redis) {
	//default setting
	opt := RedisOptions{
		MaxIdle:   100,
		MaxActive: 100,
		Timeout:   100,
		Wait:      true,
	}

	if len(cfg.Options) > 0 {
		opt = cfg.Options[0]
	}

	if cfg.Network == "" {
		cfg.Network = NetworkTCP
	}

	return &Redis{
		pool: &redigo.Pool{
			MaxIdle:     opt.MaxIdle,
			MaxActive:   opt.MaxActive,
			IdleTimeout: time.Duration(opt.Timeout) * time.Second,
			Dial: func() (redigo.Conn, error) {
				return dialer(cfg.Network, cfg.Address)
			},
		},
	}
}

// Set checks if field non empty then do hash set,
// otherwise do set
func (r *Redis) Set(key string, value string, ttl int, field ...interface{}) (interface{}, error) {
	conn := r.pool.Get()
	defer conn.Close()

	hash := len(field) > 0

	defer r.SetExpire(key, ttl)

	if hash {
		return redigo.Int64(conn.Do("HSET", key, field[0], value))
	}
	return redigo.String(conn.Do("SET", key, value))
}

//LPush into redis list
func (r *Redis) LPush(key string, value string, ttl int) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	defer r.SetExpire(key, ttl)

	return redigo.Int64(conn.Do("LPUSH", key, value))
}

//RPush into redis list
func (r *Redis) RPush(key string, value string, ttl int) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	defer r.SetExpire(key, ttl)

	return redigo.Int64(conn.Do("RPUSH", key, value))
}

//LRem Remove element from redis list
func (r *Redis) LRem(key string, value string, ttl int) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	defer r.SetExpire(key, ttl)

	return redigo.Int64(conn.Do("LREM", key, -1, value))
}

//LRange gets data from redis list
func (r *Redis) LRange(key string, limit int) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	return redigo.Strings(conn.Do("LRANGE", key, 0, limit))
}

// Get checks if field non empty then do hash get,
// otherwise do get
func (r *Redis) Get(key string, field ...interface{}) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	hash := len(field) > 0

	if hash {
		return redigo.String(conn.Do("HGET", key, field[0]))
	}
	return redigo.String(conn.Do("GET", key))
}

// Del will delete a key or a filed from cache, returning number of key deleted
func (r *Redis) Del(key string, field ...interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	hash := len(field) > 0

	if hash {
		return redigo.Int64(conn.Do("HDEL", key, field[0]))
	}
	return redigo.Int64(conn.Do("DEL", key))
}

// Exists is wrapper of redis Exists
func (r *Redis) Exists(key string) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()

	checker, err := redigo.Int64(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}

	return (checker == 1), nil
}

// AddMember add given members to the key
func (r *Redis) AddMember(key string, args []interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	return redigo.Int64(conn.Do("SADD", redigo.Args{key}.AddFlat(args)...))
}

// ReadMember read members of the given key
func (r *Redis) ReadMember(key string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	return redigo.Strings(conn.Do("SMEMBERS", key))
}

// SetExpire set expiration of a key
func (r *Redis) SetExpire(key string, ttl int) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	return redigo.Int64(conn.Do("EXPIRE", key, ttl))
}
