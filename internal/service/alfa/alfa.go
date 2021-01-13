package alfa

import "time"

type ITAlfaResource interface {
	Get(id uint64) (EmployeeModel, error)
	Add(fname string, address string, dob time.Time, role_id uint64, role_name string, salary uint64) (EmployeeModel, error)
	Edit(employee *EmployeeModel) (*EmployeeModel, error)
	Del(id uint64) (EmployeeModel, error)
	GetAll() ([]EmployeeModel, error)
	GetRole(id uint64) (string, error)
}

type Employee struct {
	employee ITAlfaResource
}

func New(t ITAlfaResource) *Employee {
	return &Employee{
		employee: t,
	}
}

//AddEmployee
func (t *Employee) AddEmployee(fname string, address string, dob time.Time, role_id uint64, role_name string, salary uint64) (EmployeeModel, error) {

	return t.employee.Add(fname, address, dob, role_id, role_name, salary)
}

//EditEmployee
func (t *Employee) EditEmployee(employee *EmployeeModel) (*EmployeeModel, error) {

	return t.employee.Edit(employee)
}

//GetID Get id employee
func (t *Employee) GetID(id uint64) (EmployeeModel, error) {

	return t.employee.Get(id)
}

//DelEmployee
func (t *Employee) DelEmployee(id uint64) (EmployeeModel, error) {

	return t.employee.Del(id)
}

// GetAllEmployee will return all employee
func (t *Employee) GetAllEmployee() ([]EmployeeModel, error) {

	return t.employee.GetAll()
}

func (t *Employee) GetRole(id uint64) (string, error) {

	return t.employee.GetRole(id)
}
