package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type citiesResponse struct {
	Cities []string `json:"cities"` // Cities capitalised to export it, otherwise json encoder will ignore it.
}

func cityHandler(res http.ResponseWriter, req *http.Request) {
	cities := citiesResponse{
		Cities: []string{"Amsterdam", "Berlin", "New York", "San Francisco", "Tokyo", "Singapore"}}

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(res).Encode(cities)
}

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.Write([]byte("Hello World!!!"))
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/cities.json", cityHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Unable to listen on port 5000 : ", err)
	}
}
