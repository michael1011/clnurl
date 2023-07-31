package router

import (
	"github.com/michael1011/clnurl/clnurl"
	"net/http"
)

func Start(cu *clnurl.ClnUrl, addr string, serveSite bool) error {
	return http.ListenAndServe(
		addr,
		registerRoutes(cu, serveSite),
	)
}
