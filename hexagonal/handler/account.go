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
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	// check json type
	if r.Header.Get("content-type") != "application/json"{
		handleError(w,errs.NewVaildationError("request body must be json"))
		return
	}

	// decode json to struct
	request := service.NewAccountRequest{}

	// check is request body is vaild json 
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		handleError(w,errs.NewVaildationError("invalid json body"))
		return
	}

	// from business logic service
	response, err := h.accSrv.NewAccount(customerID,request)
	if err != nil{
		handleError(w,err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(response)
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