package app

import "sqliteMoney/pkg/sqlite3"

type Config struct{
	BindAddr string
	Store *sqlite3.Config
}

func NewConfig() *Config{
	return &Config{
		BindAddr: "5050",
		Store: sqlite3.NewConfig(),
	}
}