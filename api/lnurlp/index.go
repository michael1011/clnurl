package handler

import (
	"encoding/json"
	"net/http"

	"github.com/michael1011/clnurl/api/_pkg/utils"
	"github.com/michael1011/clnurl/clnurl"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	backend, err := utils.GetClnBackend()
	if err != nil {
		utils.FormatError(w, http.StatusInternalServerError, err)
		return
	}

	lnurlRes, err := clnurl.Init(utils.GetConfig(), *backend).GetLnurlp()
	if err != nil {
		utils.FormatError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SetHeaders(w)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(lnurlRes)
}
