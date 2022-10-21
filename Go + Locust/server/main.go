package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	log.Println(jsonPrediction)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inicio", Inicio).Methods("GET")
	router.HandleFunc("/input", newPrediction).Methods("POST")
	log.Println("Server iniciado en el puerto 3050")
	log.Fatal(http.ListenAndServe(":3050", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
