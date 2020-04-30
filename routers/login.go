package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dekklabs/twittercopy/bd"
	"github.com/dekklabs/twittercopy/jwt"
	"github.com/dekklabs/twittercopy/models"
)

//Login realiza el inicio de sesion
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválido"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if existe == false {
		http.Error(w, "Usuario y/o contraseña inválido", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/*Como guardar una cookie*/
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
