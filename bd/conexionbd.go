package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = ConectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://diegocv:uaSr3BanfhZdWUUh@cluster0.s0vr0.mongodb.net/twiter?retryWrites=true&w=majority")

func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Printf("Conexion valida \n")
	return client
}
func CheckConnection() int {
	var result int
	err := MongoC.Ping(context.TODO(), nil)
	if err == nil {
		result = 1
	}

	return result
}
