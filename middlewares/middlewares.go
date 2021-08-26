package middlewares

import (
	"fmt"
	"net/http"
)

func AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware called")
		next.ServeHTTP(w, r)
	})

}
