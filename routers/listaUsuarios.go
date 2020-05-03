package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dekklabs/twittercopy/bd"
)

//ListaUsuarios leo la lista de los usuarios
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	typeuser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)

	result, status := bd.LeoUsariosTodos(IDusuario, pag, search, typeuser)
	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
