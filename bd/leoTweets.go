package bd

import (
	"context"
	"log"
	"time"

	"github.com/dekklabs/twittercopy/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoTweets muestra todos los tweets paginados en 20 del perfil
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweet

	//Condición en la que se va buscar en la db
	condicion := bson.M{
		"userid": ID,
	}

	//Opciones en mongodb
	opciones := options.Find()
	opciones.SetLimit(20)
	//bson.D otro formato de bson (key y valor) el -1 significa DESCENDENTE
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	//Saltar
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	//context.TODO crea un contexto vacio
	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweet
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}

		resultados = append(resultados, &registro)
	}

	return resultados, true
}
