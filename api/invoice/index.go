package handler

import (
	"github.com/michael1011/clnurl/api/_pkg/utils"
	"github.com/michael1011/clnurl/clnurl/handler"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handler.HandleInvoice(utils.GetCu, w, r)
}
