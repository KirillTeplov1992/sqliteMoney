package handlers

import (
	"encoding/json"
	"net/http"
	"sqliteMoney/internal/handlers"
	"sqliteMoney/internal/store"
)

const (
	getExpensesURL = "/get_expenses"
	getIncomsURL = "/get_incoms"
)

type handler struct{
	repository store.Store
}

func NewHandler(repository store.Store) handlers.Handler{
	return &handler{
		repository: repository,
	}
}

func (h *handler) Register(router *http.ServeMux){
	router.HandleFunc(getExpensesURL, h.getExpenses)
	router.HandleFunc(getIncomsURL, h.getIncoms)
}

func (h *handler) getExpenses(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	expList := h.repository.CategoryRep.GetExpenses()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(expList)
}

func (h *handler) getIncoms(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	incomList := h.repository.CategoryRep.GetIncoms()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(incomList)
}



