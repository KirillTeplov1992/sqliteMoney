package handlers

import (
	"encoding/json"
	"net/http"
	"sqliteMoney/internal/handlers"
	"sqliteMoney/internal/reports/models"
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

	DatesArray := []models.Dates{
		models.Dates{Month: "Весь период", Begin: "2025-10-01", End: "2026-09-30"},
		models.Dates{Month: "Октябрь 2025", Begin: "2025-10-01", End: "2025-10-31"},
		models.Dates{Month: "Ноябрь 2025", Begin: "2025-11-01", End: "2025-11-30"},
		models.Dates{Month: "Декабрь 2025", Begin: "2025-12-01", End: "2025-12-31"},
		models.Dates{Month: "Январь 2026", Begin: "2026-01-01", End: "2026-01-31"},
		models.Dates{Month: "Февраль 2026", Begin: "2026-02-01", End: "2026-02-28"},
		models.Dates{Month: "Март 2026", Begin: "2026-03-01", End: "2026-03-31"},
		models.Dates{Month: "Апрель 2026", Begin: "2026-04-01", End: "2026-04-30"},
		models.Dates{Month: "Май 2026", Begin: "2026-05-01", End: "2026-05-31"},
		models.Dates{Month: "Июнь 2026", Begin: "2026-06-01", End: "2026-06-30"},
		models.Dates{Month: "Июль 2026", Begin: "2026-07-01", End: "2026-07-31"},
	}

	var categoryReports []models.TemplCategoryReport

	for _, d := range DatesArray{
		incoms := h.repository.ReportsRep.GetIncoms(d.Begin, d.End)
		expenses := h.repository.ReportsRep.GetExpenses(d.Begin, d.End)

		tableIncoms := templates.CategoryReport(incoms)
		tableExpenses := templates.CategoryReport(expenses)

		tempReport := models.TemplCategoryReport{
			Month: d.Month,
			Incoms: tableIncoms,
			Expenses: tableExpenses,
		}

		categoryReports = append(categoryReports, tempReport)
	} 

	title := "Отчеты"

	//c это сокращенно от contents
	Plot := templates.PlotReport(balance)
	MainTable := templates.MainTable(rep, freps)
	CategoryReport := templates.MainCategoryReport(categoryReports)
	c := templates.Reports(Plot, MainTable, CategoryReport)
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
