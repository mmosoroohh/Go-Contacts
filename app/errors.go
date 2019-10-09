package app

import (
	u "github.com/mmosoroohh/go-contact/utils"
	"net/http"
)

var NotFoundHandler = func(next http.Handler) http.Handeler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "This resource was not found in the server"))
		next.ServeHTTP(w, r)
	})
}