package bd

import (
	"context"
	"time"

	"github.com/dekklabs/twittercopy/models"
	"go.mongodb.org/mongo-driver/bson"
)

//LeoTweetsSeguidores lee los tweets de mis seguidores
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relaciones")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)

	//$match
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})

	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	//Datos ordenados $sort de forma descendente
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}})
	//Saltar
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)

	var results []models.DevuelvoTweetsSeguidores

	err = cursor.All(ctx, &results)

	if err != nil {
		return results, false
	}

	return results, true
}
