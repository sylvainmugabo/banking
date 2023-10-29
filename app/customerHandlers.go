package app

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sylvainmugabo/microservices/banking/services"
)

type CustomerHandlers struct {
	service services.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	requestedStatus := r.URL.Query().Get("status")
	status := mapStatusToNumber(requestedStatus)
	cust, err := ch.service.GetAllCustomers(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())

	} else {
		writeResponse(w, http.StatusOK, cust)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cust, err := ch.service.GetCustomer(vars["customer_id"])
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, cust)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func mapStatusToNumber(status string) string {
	if strings.ToLower(status) == "active" || status == "1" {
		return "1"
	}
	if strings.ToLower(status) == "inactive" || status == "0" {
		return "0"
	}
	return ""
}
