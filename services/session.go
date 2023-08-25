package services

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var sessionSecretKey = []byte("8r4sQ#L@j6zH^W9m")
var store = sessions.NewCookieStore(sessionSecretKey)

func CreateCryptoSession(w http.ResponseWriter, r *http.Request, key string, value interface{}) error {
	session, err := store.Get(r, "crypto-session")
	if err != nil {
		return err
	}

	session.Values[key] = value
	return session.Save(r, w)
}

func GetCryptoSessionValue(r *http.Request, key string) (interface{}, error) {
	session, err := store.Get(r, "crypto-session")
	if err != nil {
		return nil, err
	}

	value := session.Values[key]
	return value, nil
}
