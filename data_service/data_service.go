package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type RespData struct {
	Time    string `json:"time,omitempty"`
	Version string `json:"version,omitempty"`
}

func GetTimeData(w http.ResponseWriter, r *http.Request) {
	log.Println("Time data requested")

	var respData []RespData
	nowTimestamp := time.Now().UTC().String()
	respData = append(respData, RespData{Time: nowTimestamp})

	json.NewEncoder(w).Encode(respData)

	log.Println("Time data delivered")
}

func GetVersionData(w http.ResponseWriter, r *http.Request) {
	log.Println("Version data requested")

	var respData []RespData
	content, err := ioutil.ReadFile("/VERSION")
	if err != nil {
		log.Println("Failed to read version file: ", err)
	}
	version := string(content)
	respData = append(respData, RespData{Version: version})

	json.NewEncoder(w).Encode(respData)

	log.Println("Version data delivered")
}

func main() {
	log.Println("Starting meshlab data service")
	router := mux.NewRouter()
	router.HandleFunc("/timedata", GetTimeData).Methods("GET")
	router.HandleFunc("/version", GetVersionData).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
