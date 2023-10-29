package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sylvainmugabo/microservices-lib/logger"
	"github.com/sylvainmugabo/microservices/banking/dto"
	"github.com/sylvainmugabo/microservices/banking/services"
)

type AccountHandler struct {
	service services.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.AccountRequest
	vars := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error("Unable to decode the request" + err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())

	} else {
		request.CustomerId = vars["customer_id"]
		account, err := h.service.NewAccount(request)
		if err != nil {
			writeResponse(w, http.StatusBadRequest, err)
		} else {

			writeResponse(w, http.StatusCreated, account)
		}

	}

}

// func (h AccountHandler) makeTransaction(w http.ResponseWriter, r *http.Request) {
// 	var request dto.TransactionRequest
// 	vars := mux.Vars(r)

// 	err := json.NewDecoder(r.Body).Decode(&request)
// 	if err != nil {
// 		logger.Error("Unable to decode the request" + err.Error())
// 		writeResponse(w, http.StatusBadRequest, err.Error())

// 	} else {

// 		request.CustomerId = vars["customer_id"]
// 		request.AccountId = vars["account_id"]
// 		account, err := h.service.MakeTransaction(request)
// 		if err != nil {
// 			writeResponse(w, http.StatusBadRequest, err)
// 		} else {

// 			writeResponse(w, http.StatusCreated, account)
// 		}

// 	}

// }

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	// decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		logger.Info("Account id is " + vars["account_id"])

		//build the request object
		request.AccountId = accountId
		request.CustomerId = customerId

		// make transaction
		account, appError := h.service.MakeTransaction(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}

}
