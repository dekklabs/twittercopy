package middlew

import (
	"net/http"

	"github.com/dekklabs/twittercopy/bd"
)

//ChequeoDB middleware de verificación a la base de datos
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
