package sqlite3

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct{
	DB *sql.DB
	config *Config
}

func NewStore(config *Config) *Store{
	return &Store{
		config: config,
	}
} 

func (s *Store) Open() error{
	db, err := sql.Open("sqlite3", s.config.DataBasePath)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil{
		return err
	}

	s.DB = db

	return nil
}

func (s *Store) Close(){
	s.DB.Close()
}