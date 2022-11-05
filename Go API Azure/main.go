package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type prediction struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Score string `json:"score"`
	Phase int    `json:"phase"`
}

var MONGO_URL = "mongodb://cosmosdb-mongo-sopes1:FQZmTPpDpwOjpqkAUdkdLrT8qnCZktOjbm3Fi6R2lddZUSWTtHre1yQGm7NEVTE5VV4zU9thxpzPACDb8hKlKA==@cosmosdb-mongo-sopes1.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@cosmosdb-mongo-sopes1@"

/*
*
Funcion para iniciar el servidor y ver que conecte
*/
func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Conexión exitosa api Go")
	log.Println("Bienvenido API")
}

func getAllMongo(w http.ResponseWriter, r *http.Request) {
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

	filter := bson.D{{}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []prediction
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func getAllRedis(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:      "azureCache-Redis.redis.cache.windows.net:6380",
		Password:  "hADq0GJgiULQECI2u1jrwRFfad9a3KgMeAzCaNfEQAU=",
		TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12},
	})
	// test connection
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	iter := client.Scan(0, "prefix:*", 1000).Iterator()
	for iter.Next() {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		fmt.Println("test errror")
		panic(err)
	}
	// return pong if server is online
	//fmt.Println(pong)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pong)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Inicio).Methods("GET")
	router.HandleFunc("/get-all-mongo", getAllMongo).Methods("GET")
	router.HandleFunc("/get-all-redis", getAllRedis).Methods("GET")
	log.Println("Server iniciado en el puerto 9093")
	log.Fatal(http.ListenAndServe(":9093", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
