package routers

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/twittercopy/bd"
	"github.com/dekklabs/twittercopy/models"
)

//ConsultoRelacion chequea si hay relación entre usuarios
func ConsultoRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	//Mi id que viene con el token || ID usuario variable global
	t.UsuarioID = IDusuario
	//Id que viene con el http || es el id del otro usuario a seguir
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)

	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
