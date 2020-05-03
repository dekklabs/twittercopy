package routers

import (
	"net/http"

	"github.com/dekklabs/twittercopy/bd"
)

//EliminarTweet elimina un tweet
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDusuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el twett"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
