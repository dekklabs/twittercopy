package routers

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/twittercopy/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadGateway)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el Registro"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(perfil)
}
