package models

//RespuestaConsultaRelacion tiene el true o el false que se obtiene al consultar la relación entre usuarios
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
