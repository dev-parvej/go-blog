package config

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(Get("SESSION_SECRET_KEY")))

func Session(r *http.Request) *sessions.Session {
	session, err := store.Get(r, "session_id")
	if err != nil {
		log.Fatal(err)
	}

	return session
}
