package middleware

import (
	"net/http"

	"github.com/dev-parvej/go-blog/config"
)

func IsAuthenticated(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var session = config.Session(r)
		authenticated := session.Values["authenticated"]

		if authenticated == nil || authenticated == false {
			http.Redirect(w, r, "/sign-in", http.StatusAccepted)
			return
		}
		h.ServeHTTP(w, r)
	})
}
