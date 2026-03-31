package handlers

import (
	"net/http"
	"sqliteMoney/internal/handlers"
	"sqliteMoney/internal/store"
	"sqliteMoney/internal/transaction/models"
	"sqliteMoney/ui/templates"
	"strconv"
	"time"
)

const (
	createTransactionURL = "/create_transaction"
	addTransactionURL = "/add_transaction"
	getTransactionURL = "/transaction"
	changeTransactionURL = "/change_transaction"
	changeTransferURL = "/change_tansfer"
)

var _ handlers.Handler = &handler{}


type handler struct{
	repository store.Store
}

func NewHandler(repository store.Store) handlers.Handler{
	return &handler{
		repository: repository,
	}
}

func (h *handler) Register(router *http.ServeMux){
	router.HandleFunc(createTransactionURL, h.createTransaction)
	router.HandleFunc(addTransactionURL, h.addTransaction)
	router.HandleFunc(getTransactionURL, h.getTransaction)
	router.HandleFunc(changeTransactionURL, h.changeTransaction)
	router.HandleFunc(changeTransferURL, h.changeTransfer)
}

func (h *handler) createTransaction(w http.ResponseWriter, r *http.Request){
	accList := h.repository.AccountRep.GetAccountsList()
	catList := h.repository.CategoryRep.GetExpenses()
	
	title := "Добавить транзакцию"

	c := templates.CreateTransaction(accList, catList)
	err := templates.Layout(c, title).Render(r.Context(), w)
	if err != nil{
		panic(err)
	}	
}

func (h *handler) addTransaction(w http.ResponseWriter, r *http.Request){
	action := r.FormValue("action")
	type_of_category := r.FormValue("type_of_category")

	date, err := time.Parse("2006-01-02" ,r.FormValue("date"))
	if err != nil {
		panic(err)
	}

	acc_id, err := strconv.Atoi(r.FormValue("account"))
	if err != nil {
		panic(err)
	}

	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil{
		panic(err)
	}

	if type_of_category != "Перевод" {
		cat_id, err := strconv.Atoi(r.FormValue("category"))
		if err != nil{
			panic(err)
		}	
			
		//var errList []string

		transaction := models.Transaction {
			Date: date,
			AccountID: acc_id,
			CategoryID: cat_id,
			Amount: amount,
			Comment: r.FormValue("comment"),
		}
	

		if action == "add" {
			h.repository.TransactionRep.AddTransaction(transaction)
			http.Redirect(w, r,"/", http.StatusSeeOther)
		}

		if action == "add+" {
			h.repository.TransactionRep.AddTransaction(transaction)
			http.Redirect(w, r, "/create_transaction", http.StatusSeeOther)
		}
	} else {
		first_acc_id := acc_id
		second_acc_id, err := strconv.Atoi(r.FormValue("category"))
		if err != nil{
			panic(err)
		}

		transactionMinus := models.Transaction {
			Date: date,
			AccountID: first_acc_id,
			CategoryID: 1,
			Amount: amount,
			Comment: r.FormValue("comment"),
		}

		transactionPlus := models.Transaction {
			Date: date,
			AccountID: second_acc_id,
			CategoryID: 2,
			Amount: amount,
			Comment: r.FormValue("comment"),
		}

		h.repository.TransactionRep.AddTransaction(transactionMinus)
		h.repository.TransactionRep.AddTransaction(transactionPlus)

		if action == "add"{
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		if action == "add+"{
			http.Redirect(w, r, "/create_transaction", http.StatusSeeOther)
		}
	}
}

func (h *handler) getTransaction (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	transaction := h.repository.TransactionRep.GetTransaction(id)
	accList := h.repository.AccountRep.GetAccountsList()


	if transaction.CategoryName == "Перевод -" || transaction.CategoryName == "Перевод +"{
		if transaction.CategoryName == "Перевод -"{
			transactionM := transaction
			transactionP := h.repository.TransactionRep.GetTransaction(id+1)
			
			c := templates.GetTransfer(transactionM, transactionP, accList)
			err := templates.Layout(c, "Меняем перевод").Render(r.Context(), w)
			if err != nil{
				panic(err)
			}
		} else {
			transactionP := transaction
			transactionM := h.repository.TransactionRep.GetTransaction(id-1)
			
			c := templates.GetTransfer(transactionM, transactionP, accList)
			err := templates.Layout(c, "Меняем перевод").Render(r.Context(), w)
			if err != nil{
				panic(err)
			}
		}
	} else{
		if transaction.TypeOfCategry == "Расход"{
			expenseList := h.repository.CategoryRep.GetExpenses()
			c := templates.GetTransaction(accList, expenseList, transaction)
			err := templates.Layout(c, "Меняем транзакцию").Render(r.Context(), w)
			if err != nil{
				panic(err)
			}
		} else {
			incomList := h.repository.CategoryRep.GetIncoms()
			c := templates.GetTransaction(accList, incomList, transaction)
			err := templates.Layout(c, "Меняем транзакцию").Render(r.Context(), w)
			if err != nil{
				panic(err)
			}
		}
	}
	
}

func (h *handler) changeTransaction (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	action := r.FormValue("action")
	
	date, err := time.Parse("2006-01-02" ,r.FormValue("date"))
	if err != nil {
		panic(err)
	}

	acc_id, err := strconv.Atoi(r.FormValue("account"))
	if err != nil {
		panic(err)
	}

	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil{
		panic(err)
	}

	cat_id, err := strconv.Atoi(r.FormValue("category"))
	if err != nil{
		panic(err)
	}

	transaction := models.Transaction{
		ID: id,
		Date: date,
		AccountID: acc_id,
		CategoryID: cat_id,
		Amount: amount,
		Comment: r.FormValue("comment"),		
	}

	h.repository.TransactionRep.UpdateTransaction(transaction)

	if action == "add"{
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if action == "add+"{
		http.Redirect(w, r, "/create_transaction", http.StatusSeeOther)
	} 
}

func (h *handler) changeTransfer (w http.ResponseWriter, r *http.Request){
	idM, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || idM < 1{
		http.NotFound(w,r)
		return
	}

	idP := idM+1
	
	date, err := time.Parse("2006-01-02" ,r.FormValue("date"))
	if err != nil {
		panic(err)
	}

	acc_idM, err := strconv.Atoi(r.FormValue("account"))
	if err != nil {
		panic(err)
	}

	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil{
		panic(err)
	}

	acc_idP, err := strconv.Atoi(r.FormValue("category"))
	if err != nil{
		panic(err)
	}

	transactionM := models.Transaction{
		ID: idM,
		Date: date,
		AccountID: acc_idM,
		CategoryID: 1,
		Amount: amount,
		Comment: r.FormValue("comment"),		
	}

	transactionP := models.Transaction{
		ID: idP,
		Date: date,
		AccountID: acc_idP,
		CategoryID: 2,
		Amount: amount,
		Comment: r.FormValue("comment"),		
	}

	h.repository.TransactionRep.UpdateTransaction(transactionM)
	h.repository.TransactionRep.UpdateTransaction(transactionP)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}


