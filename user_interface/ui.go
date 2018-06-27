package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type PageContent struct {
	Data string
}

func index(w http.ResponseWriter, r *http.Request) {

	log.Println("Index page requested")

	content := PageContent{}

	// parse page template
	tmpl, err := template.ParseFiles("/index.html")
	if err != nil {
		log.Println("Failed to parse template", err)
	}

	// call data service
	resp, err := http.Get("http://meshlab-lb-data-svc/timedata")
	if err != nil {
		log.Println("Failed to get response from load balanced data service")

		content.Data = "Nothing - it doesn't seem to be running"
	}

	if resp != nil {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Failed to read data service response")

			content.Data = "Nothing intelligble - it seems to be broken"
		}
		content.Data = string(data)
	}

	tmpl.Execute(w, content)

	log.Println("Index page served")
}

func getDataSvc() {
}

func main() {
	log.Println("Starting meshlab ui")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
