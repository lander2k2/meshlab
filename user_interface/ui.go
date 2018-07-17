package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Service struct {
	Name     string
	Endpoint string
	Response string
}

func index(w http.ResponseWriter, r *http.Request, services []Service) {

	log.Println("Index page requested")

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

func main() {
	log.Println("Starting meshlab ui")

	args := os.Args
	arg := ""
	if len(args) > 1 {
		arg = args[1]
	}

	services := []Service{
		{Name: "LB Data", Endpoint: "http://meshlab-lb-data-svc/timedata"},
		{Name: "Rate Limit Data", Endpoint: "http://meshlab-rate-limit-data-svc/timedata"},
		{Name: "Canary Data", Endpoint: "http://meshlab-canary-data-svc/version"},
	}
	if arg == "no-canary" {
		services = []Service{
			{Name: "LB Data", Endpoint: "http://meshlab-lb-data-svc/timedata"},
			{Name: "Rate Limit Data", Endpoint: "http://meshlab-rate-limit-data-svc/timedata"},
		}
	}

	//http.HandleFunc("/", index)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(w, r, services)
	})

	http.ListenAndServe(":8080", nil)
}
