package main

import (
	"alfa/internal/delivery/http/alfa"

	"github.com/gorilla/mux"
)

func newRouter(
	alfaHandler alfa.Handler,
) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/dogs", alfaHandler.GetAllDogs).Methods("GET")
	r.HandleFunc("/dogs/{breed}", alfaHandler.GetBreed).Methods("GET")

	r.HandleFunc("/employee", alfaHandler.Post).Methods("POST")
	r.HandleFunc("/employee/{id_employee}", alfaHandler.Put).Methods("PUT")
	r.HandleFunc("/employee/{id_employee}", alfaHandler.Del).Methods("DELETE")
	r.HandleFunc("/employee/{id_employee}", alfaHandler.Get).Methods("GET")
	r.HandleFunc("/employee", alfaHandler.GetAll).Methods("GET")

	return r
}
