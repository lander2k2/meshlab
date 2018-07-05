package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
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

func GetVersionDataError(w http.ResponseWriter, r *http.Request) {
	log.Println("Version data requested")

	dieRoll := rand.Int()
	if dieRoll%2 == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		log.Println("Internal server error returned")
	} else {
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
}

func GetVersionDataSlow(w http.ResponseWriter, r *http.Request) {
	log.Println("Version data requested")

	time.Sleep(5 * time.Second)

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

	args := os.Args
	arg := ""
	if len(args) > 1 {
		arg = args[1]
	}

	if arg == "error" {
		router.HandleFunc("/version", GetVersionDataError).Methods("GET")
	} else if arg == "slow" {
		router.HandleFunc("/version", GetVersionDataSlow).Methods("GET")
	} else {
		router.HandleFunc("/version", GetVersionData).Methods("GET")
	}

	router.HandleFunc("/timedata", GetTimeData).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
