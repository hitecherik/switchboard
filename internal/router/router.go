package router

import (
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/hitecherik/switchboard/internal/walker"
)

const wildcardRaw string = `\*$`

var wildcard *regexp.Regexp

func init() {
	var err error

	wildcard, err = regexp.Compile(wildcardRaw)
	if err != nil {
		panic(err)
	}
}

func BuildRouter(redirects []walker.Redirect, temporary bool) http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	code := http.StatusMovedPermanently

	if temporary {
		code = http.StatusTemporaryRedirect
	}

	for _, redirect := range redirects {
		var route *mux.Route

		if wildcard.Match([]byte(redirect.Path)) {
			stripped := redirect.Path[:len(redirect.Path)-2]
			route = r.PathPrefix(stripped)
		} else {
			route = r.Path(redirect.Path)
		}

		route.Handler(http.RedirectHandler(redirect.Destination, code))

		if redirect.Host != "" {
			route.Host(redirect.Host)
		}
	}

	return r
}
