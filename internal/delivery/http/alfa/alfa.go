package alfa

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"alfa/internal/service/alfa"
	"alfa/pkg/errors"
	"alfa/pkg/response"

	"github.com/gorilla/mux"
)

// Handler class
type Handler struct {
	thnSvc ITAlfaService
}

type ITAlfaService interface {
	GetID(id uint64) (alfa.EmployeeModel, error)
	AddEmployee(fname string, address string, dob time.Time, role_id uint64, role_name string, salary uint64) (alfa.EmployeeModel, error)
	EditEmployee(employee *alfa.EmployeeModel) (*alfa.EmployeeModel, error)
	DelEmployee(id uint64) (alfa.EmployeeModel, error)
	GetAllEmployee() ([]alfa.EmployeeModel, error)
	GetRole(id uint64) (string, error)
}

// New will create object for class Handler
func New(t ITAlfaService) Handler {
	return Handler{
		thnSvc: t,
	}
}

//Get employee
func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	vars := mux.Vars(r)
	var idEmployee string
	idEmployee = vars["id_employee"]

	u64ID, err := strconv.ParseUint(idEmployee, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}

	data, err := h.thnSvc.GetID(uint64(u64ID))
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucGet
	resp.Error.Code = http.StatusOK
	return
}

// Post employee
func (h Handler) Post(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	var (
		fname, address, dob, role_id, salary string
	)

	fname = r.PostFormValue("full_name")
	address = r.PostFormValue("address")
	dob = r.PostFormValue("dob")
	role_id = r.PostFormValue("role_id")
	salary = r.PostFormValue("salary")

	u64RoleID, err := strconv.ParseUint(role_id, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}

	u64Salary, err := strconv.ParseUint(salary, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}

	dobISO, err := time.Parse("02-01-2006", dob)
	if err != nil {
		log.Println(err)
		return
	}

	datarole, err := h.thnSvc.GetRole(u64RoleID)

	data, err := h.thnSvc.AddEmployee(fname, address, dobISO, u64RoleID, datarole, u64Salary)
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrAdd), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucAdd
	resp.Error.Code = http.StatusOK
	return
}

// Put employee
func (h Handler) Put(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	vars := mux.Vars(r)
	var (
		idEmployee, fname, address, dob string
	)

	idEmployee = vars["id_employee"]
	fname = r.PostFormValue("full_name")
	address = r.PostFormValue("salary")
	dob = r.PostFormValue("dob")

	u64IDEmployee, err := strconv.ParseUint(idEmployee, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}

	u64Salary, err := strconv.ParseUint(address, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}

	dobISO, err := time.Parse("02-01-2006", dob)
	if err != nil {
		log.Println(err)
		return
	}

	input := &alfa.EmployeeModel{
		ID:     uint64(u64IDEmployee),
		Nama:   fname,
		Salary: u64Salary,
		Dob:    dobISO,
	}

	data, err := h.thnSvc.EditEmployee(input)
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrEdit), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucEdit
	resp.Error.Code = http.StatusOK
	return
}

// Del employee
func (h Handler) Del(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	vars := mux.Vars(r)
	var idEmployee string
	idEmployee = vars["id_employee"]

	u64ID, err := strconv.ParseUint(idEmployee, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}

	data, err := h.thnSvc.DelEmployee(uint64(u64ID))
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrDel), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucDel
	resp.Error.Code = http.StatusOK
	return
}

// GetAll employee
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	data, err := h.thnSvc.GetAllEmployee()
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Code = http.StatusOK
	return
}

func (h Handler) GetAllDogs(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	response, err := http.Get("https://dog.ceo/api/breeds/list/all")

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	var tmpData alfa.DogsModel
	var mdlData []alfa.DogsModel
	for _, key := range reflect.ValueOf(result["message"]).MapKeys() {

		mydata := reflect.ValueOf(result["message"]).MapIndex(reflect.ValueOf(key.String()))
		tmpData.Breed = key.String()
		tmpData.Subr = mydata.Interface()
		mdlData = append(mdlData, tmpData)
	}

	resp.Data = mdlData
	resp.Error.Code = http.StatusOK
	return
}

func Keys(m map[string]interface{}) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (h Handler) GetBreed(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)
	vars := mux.Vars(r)
	var breed string
	breed = vars["breed"]
	responsey, err := http.Get("https://dog.ceo/api/breeds/list/all")
	response, err := http.Get("https://dog.ceo/api/breed/" + breed + "/images")

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	var resulty map[string]interface{}
	err = json.NewDecoder(responsey.Body).Decode(&resulty)
	if err != nil {
		log.Fatal(err)
	}
	var tmpData alfa.DogsDetail
	var mdlData []alfa.DogsDetail
	for _, key := range reflect.ValueOf(resulty["message"]).MapKeys() {
		if key.String() == breed {
			mydata := reflect.ValueOf(resulty["message"]).MapIndex(reflect.ValueOf(key.String()))
			var result map[string]interface{}
			err = json.NewDecoder(response.Body).Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			tmpData.Breed = key.String()
			tmpData.SubBreed = mydata.Interface()
			tmpData.Images = result["message"]
			mdlData = append(mdlData, tmpData)
		}
	}

	resp.Data = mdlData
	resp.Error.Code = http.StatusOK
	return
}
