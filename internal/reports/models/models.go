package models

import "github.com/a-h/templ"

type Report struct{
	Month string
	Incom float64
	Expense float64
	Profit float64
}


type FinalReport struct{
	Type string
	Amount float64
	Persent float64
}

type CategoryReport struct{
	Type string
	Amount float64
}

type RestMoney struct{
	Date string `json:"date"`
	Amount float64 `json:"amount"`
}

type CategoryReportByMonth struct {
	Month string
	Incoms []CategoryReport
	Expenses []CategoryReport
}

type Dates struct{
	Month string
	Begin string
	End string
}

type TemplCategoryReport struct{
	Month string
	Incoms templ.Component
	Expenses templ.Component
}