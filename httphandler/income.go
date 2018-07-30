package httphandler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"pab/model"
	"strconv"
)

func GetIncome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	i := model.GetIncome(id)
	ijson, _ := json.Marshal(i)
	fmt.Fprintf(w, string(ijson))
}

func AddIncome(w http.ResponseWriter, r *http.Request) {
	var income model.Income
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &income)
	id := model.AddIncome(income)
	idstr := strconv.Itoa(id)
	fmt.Fprintf(w, idstr)
}

func UpdateIncome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	i := model.GetIncome(id)

	if i.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := ioutil.ReadAll(r.Body)
	var ui model.Income
	json.Unmarshal(b, &ui)
	i.Source = ui.Source
	i.Amount = ui.Amount
	model.UpdateIncome(i)
	json, _ := json.Marshal(i)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func DeleteIncome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	i := model.GetIncome(id)
	if i.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	model.DeleteIncome(i)
	w.WriteHeader(http.StatusOK)
}

func GetAllIncomes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	incomes := model.GetAllIncomes()
	incmsjson, _ := json.Marshal(incomes)
	fmt.Fprintf(w, string(incmsjson))
}
