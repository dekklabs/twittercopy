package routers

import (
	"errors"
	"strings"

	"github.com/dekklabs/twittercopy/bd"
	"github.com/dekklabs/twittercopy/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Email valor de Email usando en todos los EndPoints
var Email string

//IDusuario es el ID devuelto del modelo, que se usará en todos los EndPoints
var IDusuario string

//ProcesoToken proceso token para extraer sus valores
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	//token que viene del header viene con una palábra bearer
	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDusuario = claims.ID.Hex()
		}
		return claims, encontrado, IDusuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}

	return claims, false, string(""), err
}
