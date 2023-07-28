package handler

import (
	"encoding/json"
	"net/http"
)

func HandleLnurlp(getCu getClnurl, w http.ResponseWriter, _ *http.Request) {
	cu, err := getCu(false)
	if err != nil {
		formatError(w, http.StatusInternalServerError, err)
		return
	}

	defer func() {
		cu.Disconnect()
	}()

	lnurlRes, err := cu.GetLnurlp()
	if err != nil {
		formatError(w, http.StatusInternalServerError, err)
		return
	}

	setHeaders(w)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(lnurlRes)
}
