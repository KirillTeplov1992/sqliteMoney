package handlers

import (
	"encoding/json"
	"net/http"
	"sqliteMoney/internal/handlers"
	"sqliteMoney/internal/store"
	"sqliteMoney/ui/templates"
	"strconv"
)

const (
	homeURL = "/"
	accountURL = "/account"
	getAccountsURL = "/get_accounts"
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
	router.HandleFunc(homeURL, h.home)
	router.HandleFunc(accountURL, h.account)
	router.HandleFunc(getAccountsURL, h.getAccounts)

	//подключаю CSS стили
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    router.Handle("/static/", http.StripPrefix("/static", fileServer))
}

func (h *handler) home (w http.ResponseWriter, r *http.Request){
	title := "Домашняя бухгалтерия"
	accountList := h.repository.AccountRep.GetAccounts()
	balance := h.repository.AccountRep.GetBalance()
	//c это сокращенно от contents
	c := templates.Home(balance, accountList)
	err := templates.Layout(c, title).Render(r.Context(), w)
	if err != nil{
		panic(err)
	}
}

func (h *handler) account (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	tranList := h.repository.AccountRep.GetAccountData(id)
	accName := h.repository.AccountRep.GetAccountName(id)

	c := templates.AccountPage(tranList)
	err = templates.Layout(c, accName.Name).Render(r.Context(), w)
	if err != nil{
		panic(err)
	}
}

func (h *handler) getAccounts(w http.ResponseWriter, r *http.Request){
	accList := h.repository.AccountRep.GetAccountsList()	

	w.Header().Set("account list", "application/json")

	json.NewEncoder(w).Encode(accList)
}