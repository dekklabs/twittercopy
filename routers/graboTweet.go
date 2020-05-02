package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dekklabs/twittercopy/bd"
	"github.com/dekklabs/twittercopy/models"
)

//GraboTweet permite graba el tweet en la base de datos
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDusuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
