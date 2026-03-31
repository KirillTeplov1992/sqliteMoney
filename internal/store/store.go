package store

import (
	account "sqliteMoney/internal/account/storage"
	category "sqliteMoney/internal/category/storage"
	reports "sqliteMoney/internal/reports/storage"
	transaction "sqliteMoney/internal/transaction/storage"
	"sqliteMoney/pkg/sqlite3"
)

type Store struct{
	AccountRep *account.AccountRepository
	CategoryRep *category.CategoryRepository
	ReportsRep *reports.ReportRepository
	TransactionRep *transaction.TransactionRepository
}

func NewStore(store *sqlite3.Store) *Store{
	return &Store{
		AccountRep : account.NewRepository(store),
		CategoryRep : category.NewRepository(store),
		ReportsRep: reports.NewRepository(store),
		TransactionRep: transaction.NewRepository(store),
		
	}
}

