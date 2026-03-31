package models

type Category struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Type_of_category, Is_public, Is_active bool
}