package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	urlQueryAmount = "amount"
)

func HandleInvoice(getCu getClnurl, w http.ResponseWriter, r *http.Request) {
	amount, err := strconv.ParseInt(r.URL.Query().Get(urlQueryAmount), 10, 64)
	if err != nil {
		formatError(w, http.StatusBadRequest, err)
		return
	}

	cu, err := getCu()
	if err != nil {
		formatError(w, http.StatusInternalServerError, err)
		return
	}

	invoice, err := cu.MakeInvoice(amount)
	if err != nil {
		formatError(w, http.StatusBadRequest, err)
		return
	}

	setHeaders(w)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(invoice)
}
