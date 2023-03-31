package handler

import (
	"encoding/json"
	"net/http"

	"github.com/michael1011/clnurl/api/_pkg/utils"
	"github.com/michael1011/clnurl/clnurl"
)

type invoiceResponse struct {
	Invoice string   `json:"pr"`
	Routes  []string `json:"routes"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	backend, err := utils.GetClnBackend()
	if err != nil {
		utils.FormatError(w, http.StatusInternalServerError, err)
		return
	}

	invoice, err := clnurl.Init(*backend).MakeInvoice(10, "test")
	if err != nil {
		utils.FormatError(w, http.StatusBadRequest, err)
		return
	}

	errJs, _ := json.Marshal(invoiceResponse{
		Invoice: invoice,
		Routes:  []string{},
	})

	utils.SetJsonHeader(w)
	w.WriteHeader(http.StatusOK)
	w.Write(errJs)
}
