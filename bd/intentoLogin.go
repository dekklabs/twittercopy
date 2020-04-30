package bd

import (
	"github.com/dekklabs/twittercopy/models"
	"golang.org/x/crypto/bcrypt"
)

//IntentoLogin realiza el chequeo de login a la DB
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	user, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
