package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/dekklabs/twittercopy/middlew"
	"github.com/dekklabs/twittercopy/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores seteo puerto, handler y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoDB(routers.Login)).Methods("POST")
	router.HandleFunc("/ver-perfil", middlew.ChequeoDB(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificar-perfil", middlew.ChequeoDB(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoDB(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leo-tweet", middlew.ChequeoDB(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "5000"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
