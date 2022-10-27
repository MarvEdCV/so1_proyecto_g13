package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type prediction struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Score string `json:"score"`
	Phase int    `json:"phase"`
}

/*
*
Funcion para iniciar el servidor y ver que conecte
*/
func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Conexión exitosa api Go")
	log.Println("Petición de inicio")
}

func newPrediction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var newpred prediction

	err = json.Unmarshal(body, &newpred)

	if err != nil {
		panic(err)
	}

	var jsonPrediction = string(`{"team1":"` + newpred.Team1 + `","team2":"` + newpred.Team2 + `","score":"` + newpred.Score + `","phase":` + strconv.Itoa(newpred.Phase) + `}`)
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "kafka-cluster-kafka-bootstrap:9092", "input-kafka", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	conn.WriteMessages(kafka.Message{Value: []byte(jsonPrediction)})
	log.Println(jsonPrediction)
	w.Write([]byte(jsonPrediction))
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inicio", Inicio).Methods("GET")
	router.HandleFunc("/input", newPrediction).Methods("POST")
	log.Println("Server iniciado en el puerto 80")
	log.Fatal(http.ListenAndServe(":80", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
