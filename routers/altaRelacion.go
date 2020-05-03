package routers

import (
	"net/http"

	"github.com/dekklabs/twittercopy/bd"
	"github.com/dekklabs/twittercopy/models"
)

//AltaRelacion realiza el registro  dela relacion entre usuarios
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parámetro id es necesario", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	//Mi id que viene con el token || ID usuario variable global
	t.UsuarioID = IDusuario
	//Id que viene con el http || es el id del otro usuario a seguir
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al insertar relación"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la relación"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
