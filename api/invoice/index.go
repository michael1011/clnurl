package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/michael1011/clnurl/api/_pkg/utils"
	"github.com/michael1011/clnurl/clnurl"
)

const (
	urlQueryAmount = "amount"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	amount, err := strconv.ParseInt(r.URL.Query().Get(urlQueryAmount), 10, 64)
	if err != nil {
		utils.FormatError(w, http.StatusBadRequest, err)
		return
	}

	backend, err := utils.GetClnBackend()
	if err != nil {
		utils.FormatError(w, http.StatusInternalServerError, err)
		return
	}

	invoice, err := clnurl.Init(utils.GetConfig(), *backend).MakeInvoice(amount)
	if err != nil {
		utils.FormatError(w, http.StatusBadRequest, err)
		return
	}

	utils.SetHeaders(w)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(invoice)
}
