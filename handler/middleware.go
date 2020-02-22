package handler

import (
	"log"
	"net/http"
	"os"
)

// Auth はAPIキーの検証を行います
func Auth(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			key    = os.Getenv("SECRET_KEY")
			reqKey = r.FormValue("key")
		)
		if key != reqKey {
			log.Println("invalid key")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
