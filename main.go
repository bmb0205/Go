package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TimeStruct struct {
	TimerName string `json:"timername"`
	// NewTimer  bool   `json:"newtimer"`
	// TotalTime string `json:"totaltime"`
}

type TimeStructs []TimeStruct

// json endpoint one
func returnNameAndTime(w http.ResponseWriter, r *http.Request) {
	time := TimeStructs{
		TimeStruct{TimerName: "Hello"},
	}
	fmt.Println("Endpoint hit: returnNameAndTime")
	json.NewEncoder(w).Encode(time)
}

// request handler
func handleRequests() {
	fs := http.FileServer(http.Dir("web")) // handler for web directory files
	http.Handle("/", fs)                   // registers FileServer as handler for all requests
	http.HandleFunc("/all", returnNameAndTime)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil) // launch server listening on port 8080
}

func main() {
	handleRequests()
}
