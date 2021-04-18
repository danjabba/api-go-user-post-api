// Package middlew provides a middleware to verify some information.
package middlew

import (
	tools "apirest/tools"
	"net/http"
)

// CheckDataBase checks database conection ...
func CheckDataBase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if tools.CheckConection() == false {
			http.Error(w, "conection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
