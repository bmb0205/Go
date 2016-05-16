package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type StatusStruct struct {
	TimerName string `json:"timername"`
	// TotalTime time.Time `json:"totaltime"`
	// StartTime time.Time `json:"starttime"`
	// EndTime   time.Time `json:"endtime"`
}

type StartStruct struct {
	TimerName string    `json:"timername"`
	StartTime time.Time `json:"starttime"`
	// IsNew     bool      `json:"isnew"`
}

type StopStruct struct {
	TimerName string `json:"timername"`
	// EndTime   string `json:"endtime"`
}

// func (statusStruct *StatusStruct) SetStartTime(time startTime) {
// 	statusStruct.startTime = time.Time()
// }

/*
JSON status endpoint that accepts timer information via AJAX GET request.
Returns JSON response including total accumulated time for the specified
timer and and all start/stop pairs that contributed to it.
*/
func status(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: status")

	timer := StatusStruct{
		r.URL.Query().Get("timerName"),
		// pairs of times etc
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
		fmt.Println(err)
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

	// var m map[string]string
	// m = make(map[string]string)
	// need map of timer name to start and stop values
	// accept post request of timer name and time stamp start time
	// check if timer name exists in map, set boolean to true or false
	// return total tracked time which should be current total time

	// layout := "2006-01-02T15:04:05.000Z"
	// const longForm = "Jan 2, 2006 at 3:04pm (MST)"

	r.Header.Add("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// b, err := json.Marshal(body)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println(err)
		fmt.Println(string(body))
		// fmt.Println(json.NewEncoder(w).Encode(body))
	}

	// fmt.Println(r.URL.Query().Get("startTime"))
	// t, _ := time.Parse(longForm, r.URL.Query().Get("startTime"))

	// _, isNew := m[timerName]

	// instance of TimeStruct to be used in json marshalling
	// fmt.Println(timerName)
	// timer := StartStruct{
	// timerName,
	// t,
	// isNew,
	// }

	// r.ParseForm()
	// fmt.Println(r.Form)

	// fmt.Println("lol", a.TimerName)

	// marshal timer instance, check for errors
	// b, err := json.Marshal(timer)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Check request method type, write header and handle byte version of JSON data b
	// if r.Method == "POST" {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	fmt.Println(string(b))
	// 	w.Write(b)
	// } else {
	// 	fmt.Println("Should be using a POST request...")
	// }
}

/*
JSON stop endpoint accepts timer name and time stamp via AJAX POST request.
Returns JSON response including the total tracked time for the given timer.
*/func stop(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: stop")

	// instance of TimeStruct to be used in json marshalling
	timer := StopStruct{
		r.URL.Query().Get("TimerName"),
		// r.URL.Query().Get("EndTime"),
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
		fmt.Println(err)
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
	http.ListenAndServe(":8081", nil) // launch server listening on port 8080
}

func main() {
	handleRequests()
}
