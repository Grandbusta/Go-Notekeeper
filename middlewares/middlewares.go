package middlewares

import (
	"fmt"
	"net/http"
	"notekeeper/utils"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

var excluded = make([]string, 0)

func verifyToken(r *http.Request) string {
	var x string
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		x = strings.Split(bearerToken, " ")[1]
	}
	token, err := jwt.Parse(x, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
}

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
