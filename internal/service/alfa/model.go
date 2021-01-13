package alfa

import "time"

type EmployeeModel struct {
	ID       uint64    `db:"id" json:"id"`
	Nama     string    `db:"full_name" json:"full_name"`
	Address  string    `db:"address" json:"address"`
	Dob      time.Time `db:"dob" json:"dob"`
	RoleID   uint64    `db:"role_id" json:"role_id"`
	RoleName string    `db:"role_name" json:"role_name"`
	Salary   uint64    `db:"salary" json:"salary"`
}

type DogsModel struct {
	Breed string      `json:"breed"`
	Subr  interface{} `json:"sub_breed"`
}

type DogsDetail struct {
	Breed    string      `json:"breed"`
	SubBreed interface{} `json:"sub_breed"`
	Images   interface{} `json:"images"`
}
