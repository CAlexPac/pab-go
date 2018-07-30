package main

import (
        //"fmt"
	"net/http"
        "github.com/gorilla/mux"
        "pab/httphandler"
        )

func main() {
    r := mux.NewRouter()
    

    // Income
    i := r.PathPrefix("/income").Subrouter()
    i.HandleFunc("/{id}", httphandler.GetIncome).Methods("GET")
    i.HandleFunc("/", httphandler.AddIncome).Methods("POST")
    i.HandleFunc("/{id}", httphandler.UpdateIncome).Methods("PUT")
    i.HandleFunc("/{id}", httphandler.DeleteIncome).Methods("DELETE")

    // Incomes
    ics := r.PathPrefix("/incomes").Subrouter()
    ics.HandleFunc("/", httphandler.GetAllIncomes).Methods("GET")

    http.ListenAndServe(":8000", r)

}
