package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var MONGO_URL = "mongodb://cosmosdb-mongo-sopes1:FQZmTPpDpwOjpqkAUdkdLrT8qnCZktOjbm3Fi6R2lddZUSWTtHre1yQGm7NEVTE5VV4zU9thxpzPACDb8hKlKA==@cosmosdb-mongo-sopes1.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@cosmosdb-mongo-sopes1@"

func saveGrcp(dato string) {
	url := "http://34.66.132.45:8083/client-grcp"
	method := "POST"

	payload := strings.NewReader(dato)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func main() {
	log.Println("Estoy en el consumer :)")
	conf := kafka.ReaderConfig{
		Brokers:     []string{"kafka-cluster-kafka-bootstrap:9092"},
		Topic:       "input-kafka",
		GroupID:     "g1",
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
		saveGrcp(string(message.Value))
		var doc interface{}
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
	}
}
