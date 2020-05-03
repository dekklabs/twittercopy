package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/dekklabs/twittercopy/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ConsultoRelacion consulta la relación entre 2 usuarios
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var reultado models.Relacion

	fmt.Println(reultado)

	err := col.FindOne(ctx, condicion).Decode(&reultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
