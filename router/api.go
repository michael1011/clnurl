package router

import (
	"github.com/michael1011/clnurl/clnurl"
	"github.com/michael1011/clnurl/clnurl/handler"
	"net/http"
)

func registerRoutes(cu *clnurl.ClnUrl, serveSite bool) http.Handler {
	var fileHandler http.Handler

	if serveSite {
		fileHandler = assetHandler()
	}

	getCu := func(bool) (*clnurl.ClnUrl, error) {
		return cu, nil
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
		_ = mux.Add(pattern.method, pattern.regex, pattern.handler)
	}

	return mux
}
