package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto de conexion
var MongoCN = ConectarDB()

//var clientOptions = options.Client().ApplyURI("mongodb+srv://dekk:o7IiXX85wDaSYYnp@cluster0-dyh3o.mongodb.net/twittor")

var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/twittor")

//ConectarDB devuelve la conexion a la base de datos
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa con la DB")
	return client
}

//ChequeoConnection es el ping a la base de datos
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
