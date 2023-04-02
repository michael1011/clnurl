package handler

import (
	"encoding/json"
	"net/http"
)

func HandleBech32(getCu getClnurl, w http.ResponseWriter, _ *http.Request) {
	cu, err := getCu()
	if err != nil {
		formatError(w, http.StatusInternalServerError, err)
		return
	}

	lnurl, err := cu.GetLnurlpBech32()
	if err != nil {
		formatError(w, http.StatusInternalServerError, err)
		return
	}

	setHeaders(w)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(lnurl)
}
