package cache

import (
	"testing"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
)

func initRedisTest() (*Redis, *redigomock.Conn) {
	c := redigomock.NewConn()
	dialer := func(network, addr string, opts ...redigo.DialOption) (redigo.Conn, error) {
		return c, nil
	}

	rds := newRedisWithDialer(RedisConfig{Address: "redis:6379", Network: NetworkTCP}, dialer)

	return rds, c
}

func TestRedisSet(t *testing.T) {
	rds, c := initRedisTest()
	tests := []struct {
		name   string
		cmd    string
		key    string
		value  string
		field  string
		ttl    int
		expect interface{}
	}{

		{
			name:   "Test 1 Set",
			cmd:    "SET",
			key:    "abc_124",
			value:  "abc",
			expect: "OK",
			ttl:    10,
		},
		{
			name:   "Test 2 HSet",
			cmd:    "HSET",
			key:    "abc_124",
			field:  "nilai",
			value:  "abc",
			expect: int64(1),
			ttl:    10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if tc.field != "" {
				c.Command(tc.cmd, tc.key, tc.field, tc.value).Expect(tc.expect)

				_, err := rds.Set(tc.key, tc.value, tc.ttl, tc.field)
				if err != nil {
					t.Fatal("Set Redis, error:", err)
				}
			} else {
				c.Command(tc.cmd, tc.key, tc.value).Expect(tc.expect)

				_, err := rds.Set(tc.key, tc.value, tc.ttl)
				if err != nil {
					t.Fatal("Set Redis, error:", err)
				}
			}

		})
	}
}

func TestRedisGet(t *testing.T) {
	rds, c := initRedisTest()
	tests := []struct {
		name   string
		cmd    string
		key    string
		expect interface{}
		field  string
	}{
		{
			name:   "Test 1 Get",
			cmd:    "GET",
			key:    "abc_124",
			expect: "abc_124",
		},
		{
			name:   "Test 2 HGet",
			cmd:    "HGET",
			key:    "abc_124",
			field:  "nilai",
			expect: "abc",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if tc.field != "" {
				c.Command(tc.cmd, tc.key, tc.field).Expect(tc.expect)
				_, err := rds.Get(tc.key, tc.field)
				if err != nil {
					t.Fatal("HGET Redis, error:", err)
				}
			} else {
				c.Command(tc.cmd, tc.key).Expect(tc.expect)
				_, err := rds.Get(tc.key)
				if err != nil {
					t.Fatal("GET Redis, error:", err)
				}
			}
		})
	}
}

func TestRedis_Exist(t *testing.T) {
	rds, c := initRedisTest()
	tests := []struct {
		cmd    string
		key    string
		expect int64
	}{
		{
			cmd:    "EXISTS",
			key:    "abc_124",
			expect: 1,
		},
		{
			cmd:    "EXISTS",
			key:    "abc_123",
			expect: 0,
		},
	}

	for _, tc := range tests {
		c.Command(tc.cmd, tc.key).Expect(tc.expect)

		_, err := rds.Exists(tc.key)
		if err != nil {
			t.Fatal("EXISTS Redis, error:", err)
		}
	}
}

func TestRedis_AddMember(t *testing.T) {
	rds, c := initRedisTest()
	tests := []struct {
		cmd    string
		key    string
		value  []interface{}
		expect int64
	}{
		{
			cmd:    "SADD",
			key:    "abc_124",
			value:  []interface{}{"abc", "def"},
			expect: 1,
		},
	}

	for _, tc := range tests {
		c.Command(tc.cmd, redigo.Args{tc.key}.AddFlat(tc.value)...).Expect(tc.expect)

		_, err := rds.AddMember(tc.key, tc.value)
		if err != nil {
			t.Fatal("SAdd Redis, error:", err)
		}
	}
}

func TestRedis_ReadMember(t *testing.T) {
	rds, c := initRedisTest()
	tests := []struct {
		cmd    string
		key    string
		expect []interface{}
	}{
		{
			cmd:    "SMEMBERS",
			key:    "abc_124",
			expect: []interface{}{[]byte("abc"), []byte("def")},
		},
	}

	for _, tc := range tests {
		c.Command(tc.cmd).Expect(tc.expect)

		_, err := rds.ReadMember(tc.key)
		if err != nil {
			t.Fatal("SMembers Redis, error:", err)
		}
	}
}

func TestRedis_SetExpire(t *testing.T) {
	rds, c := initRedisTest()
	tests := []struct {
		cmd    string
		key    string
		ttl    int
		expect int64
	}{
		{
			cmd:    "EXPIRE",
			key:    "abc_125",
			ttl:    3600,
			expect: 1,
		},
		{
			cmd:    "EXPIRE",
			key:    "abc_126",
			ttl:    3600,
			expect: 0,
		},
	}

	for _, tc := range tests {
		c.Command(tc.cmd, tc.key, tc.ttl).Expect(tc.expect)

		_, err := rds.SetExpire(tc.key, tc.ttl)
		if err != nil {
			t.Fatal("EXPIRE Redis, error:", err)
		}
	}
}

func TestRedis_LPush(t *testing.T) {
	rds, c := initRedisTest()
	tests := []struct {
		name   string
		cmd    string
		key    string
		value  string
		ttl    int
		expect interface{}
	}{

		{
			name:   "Test 1 Push",
			cmd:    "LPUSH",
			key:    "abc_124",
			value:  "abc",
			expect: int64(1),
			ttl:    10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.Command(tt.cmd, tt.key, tt.value).Expect(tt.expect)

			_, err := rds.LPush(tt.key, tt.value, tt.ttl)
			if err != nil {
				t.Fatal("LPUSH Redis, error:", err)
			}
		})
	}
}

func TestRedis_RPush(t *testing.T) {
	rds, c := initRedisTest()
	tests := []struct {
		name   string
		cmd    string
		key    string
		value  string
		ttl    int
		expect interface{}
	}{

		{
			name:   "Test 1 Rpush",
			cmd:    "RPUSH",
			key:    "abc_124",
			value:  "abc",
			expect: int64(1),
			ttl:    10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.Command(tt.cmd, tt.key, tt.value).Expect(tt.expect)

			_, err := rds.RPush(tt.key, tt.value, tt.ttl)
			if err != nil {
				t.Fatal("RPUSH Redis, error:", err)
			}
		})
	}
}

func TestRedis_LRem(t *testing.T) {
	rds, c := initRedisTest()
	tests := []struct {
		name   string
		cmd    string
		key    string
		value  string
		ttl    int
		expect interface{}
	}{

		{
			name:   "Test 1 Remove",
			cmd:    "LREM",
			key:    "abc_124",
			value:  "abc",
			expect: int64(1),
			ttl:    10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.Command(tt.cmd, tt.key, -1, tt.value).Expect(tt.expect)

			_, err := rds.LRem(tt.key, tt.value, tt.ttl)
			if err != nil {
				t.Fatal("Lrem Redis, error:", err)
			}
		})
	}
}

func TestRedis_LRange(t *testing.T) {
	rds, c := initRedisTest()
	tests := []struct {
		name   string
		cmd    string
		key    string
		limit  int
		expect interface{}
	}{

		{
			name:   "Test 1 Lrange",
			cmd:    "LRANGE",
			key:    "abc_124",
			limit:  2,
			expect: []interface{}{[]byte("1"), []byte("2")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.Command(tt.cmd, tt.key, 0, tt.limit).Expect(tt.expect)

			_, err := rds.LRange(tt.key, tt.limit)
			if err != nil {
				t.Fatal("Lrange Redis, error:", err)
			}
		})
	}
}
