package main

import (
	"log"

	"github.com/dekklabs/twittercopy/bd"
	"github.com/dekklabs/twittercopy/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
