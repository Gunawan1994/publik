package alfa

import (
	"log"
	"time"

	"alfa/internal/service/alfa"
	"alfa/pkg/errors"

	"github.com/jmoiron/sqlx"
)

// Resource class
type Resource struct {
	db *sqlx.DB
}

// New will return object of Resource class
func New(db *sqlx.DB) *Resource {
	return &Resource{
		db: db,
	}
}

// GetAll data
func (r *Resource) GetAll() ([]alfa.EmployeeModel, error) {
	employee := []alfa.EmployeeModel{}
	sql := `SELECT 
				id, full_name, address, dob, role_id, role_name, salary
			FROM 
				public.employee`
	err := r.db.Select(&employee, sql)
	return employee, errors.Wrap(err)
}

// Get id employee
func (r *Resource) Get(id uint64) (alfa.EmployeeModel, error) {
	sql := `SELECT 
				id, full_name, address, dob, role_id, role_name, salary
			FROM 
				public.employee 
			WHERE 
				id = $1`
	employee := alfa.EmployeeModel{}
	err := r.db.Get(&employee, sql, id)
	return employee, errors.Wrap(err)
}

// Add data
func (r *Resource) Add(fname string, address string, dob time.Time, role_id uint64, role_name string, salary uint64) (alfa.EmployeeModel, error) {
	sql := `INSERT INTO 
				public.employee 
				(full_name, address, dob, role_id, role_name, salary) 
			VALUES 
				($1, $2, $3, $4, $5, $6) 
			RETURNING 
				id, full_name, address, dob, role_id, role_name, salary`
	employee := alfa.EmployeeModel{}
	err := r.db.QueryRowx(sql, fname, address, dob, role_id, role_name, salary).StructScan(&employee)
	return employee, errors.Wrap(err)
}

// Edit data
func (r *Resource) Edit(employee *alfa.EmployeeModel) (*alfa.EmployeeModel, error) {
	sql := `UPDATE 
				public.employee 
			SET  
				full_name = :full_name, 
				salary = :salary,
				dob = :dob
			WHERE
				id= :id
			RETURNING id, full_name, address, dob, role_id, role_name, salary`

	rows, err := r.db.NamedQuery(r.db.Rebind(sql), employee)
	if err != nil {
		log.Println(err)
		return employee, errors.Wrap(err)
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.StructScan(&employee)
		return employee, errors.Wrap(err)
	}
	return employee, errors.Wrap(err)

}

// Del data
func (r *Resource) Del(id uint64) (alfa.EmployeeModel, error) {
	sql := `DELETE
			FROM 
				public.employee
			WHERE 
				id=$1
			RETURNING 
				id, full_name, address, dob, role_id, role_name, salary`
	employee := alfa.EmployeeModel{}
	err := r.db.QueryRowx(sql, id).StructScan(&employee)
	return employee, errors.Wrap(err)
}

func (r *Resource) GetRole(id uint64) (string, error) {
	sql := `SELECT 
				name_role
			FROM 
				public.role 
			WHERE 
				id = $1`
	var role string
	err := r.db.Get(&role, sql, id)
	return role, errors.Wrap(err)
}
