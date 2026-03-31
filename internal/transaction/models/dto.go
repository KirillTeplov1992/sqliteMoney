package models

import "time"

type GottenTransaction struct {
	ID int
	Date time.Time
	AccountId int
	CategoryID int
	TypeOfCategry string
	CategoryName string
	Amount float64
	Comment string
}