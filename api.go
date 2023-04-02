package main

import (
	"fmt"
	"github.com/elementsproject/glightning/glightning"
	"github.com/michael1011/clnurl/clnurl"
	"github.com/michael1011/clnurl/clnurl/handler"
	"net/http"
)

var getCu = func() (*clnurl.ClnUrl, error) {
	return cu, nil
}

func registerRoutes(cfg *config) http.Handler {
	var fileHandler http.Handler

	if cfg.ServeSite {
		fileHandler = assetHandler()
	}

	mux := InitRegexMux(fileHandler)

	for _, pattern := range []struct {
		method  string
		regex   string
		handler http.HandlerFunc
	}{
		{
			method: http.MethodGet,
			regex:  "/api/invoice",
			handler: func(w http.ResponseWriter, r *http.Request) {
				handler.HandleInvoice(getCu, w, r)
			},
		},
		{
			method: http.MethodGet,
			regex:  "/api/lnurlp/bech32",
			handler: func(w http.ResponseWriter, r *http.Request) {
				handler.HandleBech32(getCu, w, r)
			},
		},
		{
			method: http.MethodGet,
			regex:  "/api/lnurlp",
			handler: func(w http.ResponseWriter, r *http.Request) {
				handler.HandleLnurlp(getCu, w, r)
			},
		},
		{
			method: http.MethodGet,
			regex:  "/.well-known/lnurlp/.*",
			handler: func(w http.ResponseWriter, r *http.Request) {
				handler.HandleLnurlp(getCu, w, r)
			},
		},
	} {
		err := mux.Add(pattern.method, pattern.regex, pattern.handler)
		if err != nil {
			plugin.Log(fmt.Sprintf(
				"Could not register route %s: %s",
				pattern.regex,
				err.Error(),
			), glightning.Info)
		}
	}

	return mux
}
