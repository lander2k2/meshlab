package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type RespDatum struct {
	Time string `json:"time,omitempty"`
}

func GetData(w http.ResponseWriter, r *http.Request) {
	log.Println("Time data requested")

	var respData []RespDatum
	nowTimestamp := time.Now().UTC().String()
	respData = append(respData, RespDatum{Time: nowTimestamp})

	json.NewEncoder(w).Encode(respData)

	log.Println("Time data delivered")
}

func main() {
	log.Println("Starting meshlab data service")
	router := mux.NewRouter()
	router.HandleFunc("/timedata", GetData).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
