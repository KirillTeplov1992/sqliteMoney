package app

import (
	"fmt"
	"net/http"
	account "sqliteMoney/internal/account/handlers"
	category "sqliteMoney/internal/category/handlers"
	reports "sqliteMoney/internal/reports/handlers"
	"sqliteMoney/internal/store"
	transaction "sqliteMoney/internal/transaction/handlers"
	"sqliteMoney/pkg/sqlite3"
)

type Application struct{
	router *http.ServeMux
	store *store.Store
	config *Config
}

func New() *Application{
	return &Application{
		router: http.NewServeMux(),
		config: NewConfig(),
	}
}

func (app *Application) Start() error{
	fmt.Println("Запуск веб-сервера на http://127.0.0.1:5050")

	if err := app.configureStore(); err != nil{
		return err
	}

	accHandler := account.NewHandler(*app.store)
	accHandler.Register(app.router)

	tranHandler := transaction.NewHandler(*app.store)
	tranHandler.Register(app.router)

	catHamdler := category.NewHandler(*app.store)
	catHamdler.Register(app.router)

	reportHandler := reports.NewHandler(*app.store)
	reportHandler.Register(app.router)


	return http.ListenAndServe(":5050", app.router)
}

func (app *Application) configureStore() error{
	st := sqlite3.NewStore(app.config.Store)
	if err := st.Open(); err != nil{
		return err
	}

	store := store.NewStore(st)

	app.store = store

	return nil
}