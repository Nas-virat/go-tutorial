package handler

import (
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct{
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler{
	return accountHandler{accSrv:accSrv}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request){
	
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request){
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	responses,err := h.accSrv.GetAccounts(customerID)
	if err != nil{
		handleError(w,err)
	}
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(responses)
}