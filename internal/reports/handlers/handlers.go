package handlers

import (
	"encoding/json"
	"net/http"
	"sqliteMoney/internal/handlers"
	"sqliteMoney/internal/store"
	"sqliteMoney/ui/templates"
)

const (
	reportsURL = "/reports"
	pointsURL = "/points"
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
	router.HandleFunc(reportsURL, h.getReports)
	router.HandleFunc(pointsURL, h.getRestOfTheMoney)
}

func (h *handler) getReports(w http.ResponseWriter, r *http.Request){
	balance := h.repository.AccountRep.GetBalance()
	rep := h.repository.ReportsRep.Report("2025-10-01", "2026-09-30")
	freps := h.repository.ReportsRep.GetFinalReport("2025-10-01", "2026-09-30")
	cirep := h.repository.ReportsRep.GetIncoms("2025-10-01", "2026-09-30")
	cerep := h.repository.ReportsRep.GetExpenses("2025-10-01", "2026-09-30")

	title := "Отчеты"

	//c это сокращенно от contents
	c := templates.Reports(balance, rep, freps, cirep, cerep)
	err := templates.Layout(c, title).Render(r.Context(), w)
	if err != nil{
		panic(err)
	}
}

func (h *handler) getRestOfTheMoney(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	listOfThePoints := h.repository.ReportsRep.GetPlotInfomation()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(listOfThePoints)
}
