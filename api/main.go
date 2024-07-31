package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
)

type Car struct {
	Name            string `json:"name"`
	Price           int    `json:"price"`
	Image           string `json:"image"`
	AosDelay        int    `json:"aosDelay"`
	Description     string `json:"description"`
	SeatingCapacity string `json:"seatingCapacity"`
}

type Testimonial struct {
	Name        string `json:"name"`
	Stars       string `json:"stars"`
	Image       string `json:"image"`
	Description string `json:"description"`
	AosDelay    string `json:"aosDelay"`
}

var validTokens = []string{"token1", "token2", "token3"} // replace with your actual tokens

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		jsonFile, err := ioutil.ReadFile("db.json")
		if err != nil {
			log.Printf("Error reading JSON file: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var cars []Car
		err = json.Unmarshal(jsonFile, &cars)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(cars)
		if err != nil {
			log.Printf("Error encoding JSON: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Served cars data to %s", r.RemoteAddr)
	})

	// Testimonials endpoint
	mux.HandleFunc("/testimonials", func(w http.ResponseWriter, r *http.Request) {
		jsonFile, err := ioutil.ReadFile("tes.json")
		if err != nil {
			log.Printf("Error reading JSON file: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var testimonials []Testimonial
		err = json.Unmarshal(jsonFile, &testimonials)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(testimonials)
		if err != nil {
			log.Printf("Error encoding JSON: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Served testimonials data to %s", r.RemoteAddr)
	})

	handler := cors.Default().Handler(mux)

	serverAddress := ":8080"
	fmt.Printf("Starting server on %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, handler))
}
