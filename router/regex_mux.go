package router

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type RegexMux struct {
	handlers    []regexRoute
	fileHandler http.Handler
}

type regexRoute struct {
	method  string
	pattern *regexp.Regexp
	handler http.HandlerFunc
}

func InitRegexMux(fileHandler http.Handler) *RegexMux {
	return &RegexMux{
		fileHandler: fileHandler,
	}
}

func (r *RegexMux) Add(method, regex string, handler http.HandlerFunc) error {
	compiled, err := regexp.Compile(regex)
	if err != nil {
		return fmt.Errorf("could not compiled regex: %s", err.Error())
	}

	r.handlers = append(r.handlers, regexRoute{
		method:  method,
		pattern: compiled,
		handler: handler,
	})
	return nil
}

func (r *RegexMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.handlers {
		if req.Method != route.method {
			continue
		}

		if route.pattern.MatchString(strings.ToLower(req.URL.Path)) {
			route.handler(w, req)
			return
		}
	}

	if r.fileHandler != nil {
		r.fileHandler.ServeHTTP(w, req)
	} else {
		http.NotFound(w, req)
	}
}
