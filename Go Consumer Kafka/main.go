package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

var MONGO_URL = "mongodb://cosmosdb-mongo-sopes1:FQZmTPpDpwOjpqkAUdkdLrT8qnCZktOjbm3Fi6R2lddZUSWTtHre1yQGm7NEVTE5VV4zU9thxpzPACDb8hKlKA==@cosmosdb-mongo-sopes1.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@cosmosdb-mongo-sopes1@"

func main() {
	log.Println("Estoy en el consumer :)")
	/*
		conn, _ := kafka.DialLeader(context.Background(), "tcp", "kafka-cluster-kafka-bootstrap:9092", "input-kafka", 0)
		conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

		message, _ := conn.ReadMessage(1e6)
		/*bytes := make([]byte, 1e3)
		for {
			_, err := batch.Read(bytes)
			if err != nil {
				break
			}
		fmt.Println(string(message.Value))
		//}
	*/
	conf := kafka.ReaderConfig{
		Brokers:     []string{"kafka-cluster-kafka-bootstrap:9092"},
		Topic:       "input-kafka",
		GroupID:     "matches-group",
		StartOffset: kafka.LastOffset,
	}
	reader := kafka.NewReader(conf)

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Ocurrio un error", err)
			continue
		}
		fmt.Println("El mensaje es: ", string(message.Value))
	}
	/*var doc interface{}
	errr := bson.UnmarshalExtJSON([]byte(message.Value), true, &doc)
	if errr != nil {
		log.Fatal(errr)
	}
	//CONEXION A LA BASE DE DATOS E INSERCION DE DATOS
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("UsactarMongoDB").Collection("data")
	res, insertErr := collection.InsertOne(ctx, doc)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)
	*/
}
