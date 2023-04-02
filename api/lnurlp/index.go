package handler

import (
	"github.com/michael1011/clnurl/clnurl/handler"
	"net/http"

	"github.com/michael1011/clnurl/api/_pkg/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handler.HandleLnurlp(utils.GetCu, w, r)
}
