package models


import "time"

type Transaction struct {
	ID int          
	Date time.Time 
	AccountID int   
	CategoryID int  
	Amount float64 
	Comment string   
}
