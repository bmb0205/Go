package main

import (
	"encoding/json"
	"fmt"
	// "io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type StatusStruct struct {
	TimerName string        `json:"timername"`
	TotalTime time.Duration `json:"totaltime"`
	// TimePairs map[string]time.Time `json:"timepairs"`
}

type StartStruct struct {
	TimerName string    `json:"timername"`
	StartTime time.Time `json:"starttime"`
	IsNew     bool      `json:"isnew"`
}

type StopStruct struct {
	TimerName string    `json:"timername"`
	StopTime  time.Time `json:"stoptime"`
}

// JSON status endpoint that accepts timer information via AJAX GET request.
// Returns JSON response including total accumulated time for the specified
// timer and and all start/stop pairs that contributed to it.
func status(w http.ResponseWriter, r *http.Request, myMap map[string][]map[string]time.Time) {
	fmt.Println(" ")
	fmt.Println("Endpoint hit: status")

	timerName := r.URL.Query().Get("timerName")

	var totalTime time.Duration

	for _, valueMap := range myMap {
		var tot time.Duration
		for i := range valueMap {
			start := valueMap[i]["startTime"]
			stop := valueMap[i]["stopTime"]
			tot += stop.Sub(start)
		}
		totalTime = (tot / time.Second) //(stop.Sub(start) / time.Second)

	}
	timer := StatusStruct{
		timerName,
		totalTime,
		// pairs of start/stop times
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
		w.Write(b)
	} else {
		fmt.Println("Should be using a GET request...")
	}

}

// JSON start endpoint accepts timer name and time stamp via AJAX POST request.
// Returns JSON response including the total tracked time for the given timer and
// a created boolean indicating if the timer is new.
func start(w http.ResponseWriter, r *http.Request, myMap map[string][]map[string]time.Time) {
	fmt.Println(" ")
	fmt.Println("Endpoint hit: start")

	// TODO: need map of timer name to start and stop values
	// accept post request of timer name and time stamp start time
	// check if timer name exists in map, set boolean to true or false
	// return total tracked time which should be current total time

	r.Header.Add("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// unmarshals byte stream of json string request into StartStruct instance
	// writes request body back as response...change this to give back total time
	if r.Method == "POST" {
		startTime := time.Now()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(body)
		bytes := []byte(body)
		var s StartStruct
		json.Unmarshal(bytes, &s)

		subMap := map[string]time.Time{
			"startTime": startTime,
			"stopTime":  time.Now(), // placeholder for future real stop time
		}
		myMap[s.TimerName] = append(myMap[s.TimerName], subMap)

	} else {
		fmt.Errorf("should be using a POST request...")
	}
}

// JSON stop endpoint accepts timer name and time stamp via AJAX POST request.
// Returns JSON response including the total tracked time for the given timer.
func stop(w http.ResponseWriter, r *http.Request, myMap map[string][]map[string]time.Time) {
	fmt.Println(" ")
	fmt.Println("Endpoint hit: stop")

	r.Header.Add("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// unmarshals byte stream of json string request into StopStruct instance,
	// writes request body back as response...change this to give back total time
	if r.Method == "POST" {
		stopTime := time.Now() // this needs the delay subtracted from it for a true value
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(body)
		bytes := []byte(body)
		var s StopStruct
		json.Unmarshal(bytes, &s)
		myMap[s.TimerName][len(myMap[s.TimerName])-1]["stopTime"] = stopTime
		fmt.Println(myMap)
	} else {
		fmt.Errorf("should be using a POST request...")
	}
}

// Request handler
func handleRequests() {
	myMap := make(map[string][]map[string]time.Time)
	fs := http.FileServer(http.Dir("web")) // handler for web directory files
	http.Handle("/", fs)                   // registers FileServer as handler for all requests
	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) { start(w, r, myMap) })
	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) { stop(w, r, myMap) })
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) { status(w, r, myMap) })
	log.Println("Listening...")
	http.ListenAndServe(":8081", nil) // launch server listening on port 8081
}

func main() {
	handleRequests()
}
