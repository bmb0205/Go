package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TimeStruct struct {
	TimerName   string `json:"timername"`
	StartTime   string `json:"starttime"`
	EndTime     string `json:"endtime"`
	ElapsedTime string `json:"elapsedTime"`
}

// json endpoint one
func status(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: status")

	timerName := r.URL.Query().Get("timerName")
	startTime := r.URL.Query().Get("startTime")
	endTime := r.URL.Query().Get("endTime")
	elapsedTime := r.URL.Query().Get("elapsedTime")
	fmt.Println("timer name is: ", timerName)
	fmt.Println("start time is: ", startTime)
	fmt.Println("end time is: ", endTime)
	fmt.Println("elapsed time is: ", elapsedTime)

	// var myMap map[string]string
	// myMap = make(map[string]string)

	// myMap["startTime"] = r.URL.Query().Get("startTime")
	// fmt.Println(myMap)

	timer := TimeStruct{timerName, startTime, endTime, elapsedTime}
	b, err := json.Marshal(timer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(b)
	}
}

// endpoint two
func start(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: start")

}

// endpoint three
func stop(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: stop")

}

// request handler
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
