package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		fmt.Println("Middleware called")
		next.ServeHTTP(w, r)
	})

}
