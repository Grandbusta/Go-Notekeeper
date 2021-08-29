package middlewares

import (
	"fmt"
	"net/http"
	"notekeeper/utils"

	"github.com/gorilla/mux"
)

var excluded = make([]string, 0)

func AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(excluded) == 0 {
			fmt.Println("Middleware called normally")
		} else {
			if !utils.Contains(excluded, r.URL.Path) {
				fmt.Println("Middleware called ")
			}
		}
		next.ServeHTTP(w, r)
	})
}

func Auth(m *mux.Router, exclude []string) {
	excluded = exclude
	m.Use(AuthUser)
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		if r.Method == http.MethodOptions {
			return
		}
		fmt.Println("cors handled")
		next.ServeHTTP(w, r)
	})

}
