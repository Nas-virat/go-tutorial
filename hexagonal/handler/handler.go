package handler

import (
	"bank/errs"
	"fmt"
	"net/http"
)

func handleError(w http.ResponseWriter, err error){
	// check data type
	switch e := err.(type){
		case errs.AppError:
			w.WriteHeader(e.Code)
			fmt.Fprintln(w,e)
		case error:
			w.WriteHeader(http.StatusInternalServerError)
	}
}