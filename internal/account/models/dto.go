package models

import "time"

type AccountData struct{
	ID int
	Name string
	Amount float32
}

type TotalBalance struct{
	Balance float32
}

type AccountTransaction struct{
	ID int
	Date time.Time
	Category string
	Amount float32
	Comment string
	TypeOfCategory string
}