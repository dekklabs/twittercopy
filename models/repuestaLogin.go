package models

//RespuestaLogin tiene el token que devuelve con el login
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
