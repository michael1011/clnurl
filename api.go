package main

import (
	"github.com/michael1011/clnurl/clnurl"
	"github.com/michael1011/clnurl/clnurl/handler"
	"net/http"
)

var getCu = func() (*clnurl.ClnUrl, error) {
	return cu, nil
}

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/invoice", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleInvoice(getCu, w, r)
	})
	mux.HandleFunc("/api/lnurlp", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleLnurlp(getCu, w, r)
	})
	mux.HandleFunc("/api/lnurlp/bech32", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleBech32(getCu, w, r)
	})
}
