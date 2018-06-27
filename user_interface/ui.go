package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Service struct {
	Name     string
	Endpoint string
	Response string
}

func index(w http.ResponseWriter, r *http.Request) {

	log.Println("Index page requested")

	services := []Service{
		{Name: "LB Data", Endpoint: "http://meshlab-lb-data-svc/timedata"},
		{Name: "Rate Limit Data", Endpoint: "http://meshlab-rate-limit-data-svc/timedata"},
	}

	tmpl, err := template.ParseFiles("/index.html")
	if err != nil {
		log.Println("Failed to parse template", err)
	}

	for i, svc := range services {
		resp, err := http.Get(svc.Endpoint)
		if err != nil {
			log.Println("Failed to get response from load balanced data service")
			services[i].Response = "Nothing - it doesn't seem to be running"
		}

		if resp != nil {
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("Failed to read data service response")
				services[i].Response = "Nothing intelligble - it seems to be broken"
			}
			services[i].Response = string(data)
		}
	}

	tmpl.Execute(w, services)

	log.Println("Index page served")
}

func getDataSvc() {
}

func main() {
	log.Println("Starting meshlab ui")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
