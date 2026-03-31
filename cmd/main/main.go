package main

import "sqliteMoney/internal/app"

func main(){
	newApp := app.New()

	newApp.Start()
}