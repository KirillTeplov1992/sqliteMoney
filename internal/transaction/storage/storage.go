package storage

import (
	"sqliteMoney/internal/transaction/models"
	"sqliteMoney/pkg/sqlite3"
)


type TransactionRepository struct{
	store *sqlite3.Store
}

func NewRepository(store *sqlite3.Store) *TransactionRepository{
	return &TransactionRepository{
		store: store,
	}
}

func (tr *TransactionRepository) AddTransaction(transaction models.Transaction){
	stmt := `
	INSERT INTO transactions (
		date,
		account_id,
		category_id,
		amount,
		comment)
	VALUES (
		?,
		?,
		?,
		?,
		?)
	`
	_, err := tr.store.DB.Exec(stmt,
									transaction.Date,
									transaction.AccountID,
									transaction.CategoryID,
									transaction.Amount,
									transaction.Comment)
	if err != nil{
		panic(err)
	}
}

func (tr *TransactionRepository) GetTransaction(id int) *models.GottenTransaction {
	stmt :=`
	SELECT
		T.id
		,T.date
		,T.account_id
		,T.category_id
		,C.type_of_category
		,C.name
		,T.amount
		,T.comment
	FROM
		transactions AS T
	INNER JOIN
		categories AS C
		ON
			T.category_id = C.id
	WHERE
		T.id = ?`

	row := tr.store.DB.QueryRow(stmt, id)

	gTran := &models.GottenTransaction{}

	err := row.Scan(&gTran.ID,
					&gTran.Date,
					&gTran.AccountId,
					&gTran.CategoryID,
					&gTran.TypeOfCategry,
					&gTran.CategoryName,
					&gTran.Amount,
					&gTran.Comment)
	if err != nil{
		panic(err)
	}

	return gTran
}

func (tr *TransactionRepository) UpdateTransaction (transaction models.Transaction){
	stmt :=`
	UPDATE
		transactions
	SET
		date = ?,
		account_id = ?,
		category_id = ?,
		amount = ?,
		comment = ?
	WHERE
		id = ?`

	_, err := tr.store.DB.Exec(stmt,
		 					transaction.Date,
							transaction.AccountID,
							transaction.CategoryID,
							transaction.Amount,
							transaction.Comment,
							transaction.ID)
	if err != nil{
		panic(err)
	}
}
