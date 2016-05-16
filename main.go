package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StatusStruct struct {
	TimerName string `json:"timername"`
	StartTime string `json:"starttime"`
	EndTime   string `json:"endtime"`
}

type StartStruct struct {
	TimerName string `json:"timername"`
	StartTime string `json:"starttime"`
}

type StopStruct struct {
	TimerName string `json:"timername"`
	EndTime   string `json:"endtime"`
}

/*
JSON status endpoint that accepts timer information via AJAX GET request.
Returns JSON response including total accumulated time for the specified
timer and and all start/stop pairs that contributed to it.
*/
func status(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: status")

	timer := StatusStruct{
		r.URL.Query().Get("timerName"),
		r.URL.Query().Get("startTime"),
		r.URL.Query().Get("endTime"),
	}

	// marshal timer instance, check for errors
	b, err := json.Marshal(timer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check request method type, write header and handle byte version of JSON data b
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		fmt.Println(string(b))
		w.Write(b)
	} else {
		fmt.Println("Should be using a GET request...")
	}
}

/*
JSON start endpoint accepts timer name and time stamp via AJAX POST request.
Returns JSON response including the total tracked time for the given timer and
a created boolean indicating if the timer is new.
*/
func start(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: start")

	// instance of TimeStruct to be used in json marshalling
	timer := StartStruct{
		r.URL.Query().Get("timerName"),
		r.URL.Query().Get("startTime"),
	}

	// marshal timer instance, check for errors
	b, err := json.Marshal(timer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check request method type, write header and handle byte version of JSON data b
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		fmt.Println(string(b))
		w.Write(b)
	} else {
		fmt.Println("Should be using a POST request...")
	}
}

/*
JSON stop endpoint accepts timer name and time stamp via AJAX POST request.
Returns JSON response including the total tracked time for the given timer.
*/func stop(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: stop")

	// instance of TimeStruct to be used in json marshalling
	timer := StopStruct{
		r.URL.Query().Get("TimerName"),
		r.URL.Query().Get("EndTime"),
	}

	// marshal timer instance, check for errors
	b, err := json.Marshal(timer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check request method type, write header and handle byte version of JSON data b
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		fmt.Println(string(b))
		w.Write(b)
	} else {
		fmt.Println("Should be using a POST request...")
	}
}

// Request handler
func handleRequests() {
	fs := http.FileServer(http.Dir("web")) // handler for web directory files
	http.Handle("/", fs)                   // registers FileServer as handler for all requests
	http.HandleFunc("/status", status)
	http.HandleFunc("/start", start)
	http.HandleFunc("/stop", stop)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil) // launch server listening on port 8080
}

func main() {
	handleRequests()
}
