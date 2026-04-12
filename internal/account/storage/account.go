package storage

import (
	"sqliteMoney/internal/account/models"
	"sqliteMoney/pkg/sqlite3"
)

type AccountRepository struct{
	store *sqlite3.Store
}

func NewRepository(store *sqlite3.Store) *AccountRepository{
	return &AccountRepository{
		store: store,
	}
}

func (ar *AccountRepository) GetAccounts() []*models.AccountData{
	stmt := `
	WITH I AS (
	SELECT
		A.id
		, A.name
		, sum(amount) AS amount
	FROM
		transactions AS T
	INNER JOIN
		accounts AS A
		ON
			T.account_id = A.id
	INNER JOIN
		categories AS C
		ON
			T.category_id = C.id
	WHERE
		type_of_category = "Доход"
	GROUP BY
		A.name
	ORDER BY
		A.id),
	E AS (
	SELECT
		A.id
		, A.name
		, sum(amount) AS amount
	FROM
		transactions AS T
	INNER JOIN
		accounts AS A
		ON
			T.account_id = A.id
	INNER JOIN
		categories AS C
		ON
			T.category_id = C.id
	WHERE
		type_of_category = "Расход"
	GROUP BY
		A.name
	ORDER BY
		A.id)
	
	SELECT
		I.id
		, I.name
		, ROUND(I.amount - E.amount, 2) AS amount
	FROM
		I
	INNER JOIN
		E
		ON
			I.id=E.id`

	res, err := ar.store.DB.Query(stmt)
	if err != nil{
		panic(err)
	}
	
	var accountList []*models.AccountData

	for res.Next(){
		account := &models.AccountData{}
		err = res.Scan(&account.ID,
						&account.Name,
						&account.Amount)
		if err != nil{
			panic(err)
		}

		accountList = append(accountList, account)
	}

	return accountList
}

func (ar *AccountRepository) GetBalance() *models.TotalBalance{
		stmt := `
		WITH I AS (
			SELECT 
				SUM(amount) as Incom
			FROM 
				transactions AS T
			INNER JOIN
				categories AS C
				ON
					T.category_id = C.id
			WHERE
				type_of_category == "Доход"),
		E AS (
			SELECT 
				SUM(amount) as Expenses
			FROM 
				transactions AS T
			INNER JOIN
				categories AS C
				ON
					T.category_id = C.id
			WHERE
				type_of_category == "Расход")
	
		SELECT 
			I.Incom-E.Expenses AS Profit
		FROM
			I, E`

	row := ar.store.DB.QueryRow(stmt, nil)

	b := &models.TotalBalance{}

	err := row.Scan(&b.Balance)
	if err != nil{
		panic(err)
	}

	return b
}

func (ar *AccountRepository) GetAccountData (acc_id int) []*models.AccountTransaction{
	stmt := `
	SELECT
    	T.id,
    	T.date,
    	C.name,
    	T.amount,
		T.comment,
		C.type_of_category
	FROM
		transactions T
	INNER JOIN
		categories C
	ON
		T.category_id = C.id
	WHERE
		account_id = ?
	ORDER BY
		T.date DESC
	`
	res, err := ar.store.DB.Query(stmt, acc_id)
	if err != nil{
		panic(err)
	}

	var accDataList []*models.AccountTransaction

	for res.Next(){
		transaction := &models.AccountTransaction{}
		err = res.Scan(&transaction.ID,
						&transaction.Date,
						&transaction.Category,
						&transaction.Amount,
						&transaction.Comment,
						&transaction.TypeOfCategory)
		if err != nil{
			panic(err)
		}

		accDataList = append(accDataList, transaction)
	}

	return accDataList
}

func (ar *AccountRepository) GetAccountName(acc_id int) *models.Account{
	stmt :=`
	SELECT 
		name
	FROM
		accounts
	WHERE
		id = ?
	`
	res := ar.store.DB.QueryRow(stmt, acc_id)
	
	account := &models.Account{}

	err := res.Scan(&account.Name)
	if err != nil{
		panic(err)
	}

	return account
}

func (ar *AccountRepository) GetAccountsList() []*models.Account{
	stmt := `
	SELECT
		id,
		name
	FROM
		accounts`

	res, err := ar.store.DB.Query(stmt)
	if err != nil{
		panic(err)
	}

	var accList []*models.Account

	for res.Next(){
		acc := &models.Account{}
		err = res.Scan(&acc.ID, &acc.Name)
		if err != nil{
			panic(err)
		}

		accList = append(accList, acc)
	}

	return accList

}