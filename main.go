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
	TimerName string `json:"timername"`
	// TotalTime time.Time   `json:"totaltime"`
	// TimePairs []time.Time `json:"timepairs"`
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

/*
JSON status endpoint that accepts timer information via AJAX GET request.
Returns JSON response including total accumulated time for the specified
timer and and all start/stop pairs that contributed to it.
*/
func status(w http.ResponseWriter, r *http.Request, myMap map[string][]map[string]time.Time) {
	fmt.Println("Endpoint hit: status")

	timerName := r.URL.Query().Get("timerName")
	fmt.Println("lol", timerName)
	timer := StatusStruct{
		timerName,
		// totalTime,
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
		fmt.Println(err)
		bytes := []byte(b)
		fmt.Println(string(bytes))

		w.Write(b)
	} else {
		fmt.Println("Should be using a GET request...")
	}

	// _, ok := myMap[timerName]
	// fmt.Println(ok)
	for k, v := range myMap {
		fmt.Println("K: ", k)
		fmt.Println("V: ", v)
		fmt.Println(" ")
		w.Write([]byte(k))
	}
	// fmt.Println(myMap)
}

/*
JSON start endpoint accepts timer name and time stamp via AJAX POST request.
Returns JSON response including the total tracked time for the given timer and
a created boolean indicating if the timer is new.
*/
func start(w http.ResponseWriter, r *http.Request, myMap map[string][]map[string]time.Time) {
	fmt.Println("Endpoint hit: start")
	// start := time.Now()
	// var m map[string]time.Time
	// m = make(map[string]time.Time)
	// need map of timer name to start and stop values
	// accept post request of timer name and time stamp start time
	// check if timer name exists in map, set boolean to true or false
	// return total tracked time which should be current total time

	// s := New()
	// s.start()

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

		// s.StartTime is time stamp sent from js ajax post request
		// might be 2 sec delay from s.StartTime and startTime (time.Now())
		// true start time is s.StartTime

		subMap := map[string]time.Time{
			"startTime": startTime,
			"stopTime":  time.Now(), // placeholder for future real stop time
		}
		// fmt.Println(startTime)
		// fmt.Println(s.StartTime)
		// fmt.Println(startTime.Sub(s.StartTime))
		// fmt.Println(s.StartTime)
		myMap[s.TimerName] = append(myMap[s.TimerName], subMap)
		// fmt.Println(myMap)

		// fmt.Println("endddddddd")

	} else {
		fmt.Errorf("should be using a POST request...")
	}
}

/*
JSON stop endpoint accepts timer name and time stamp via AJAX POST request.
Returns JSON response including the total tracked time for the given timer.
*/
func stop(w http.ResponseWriter, r *http.Request, myMap map[string][]map[string]time.Time) {
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

		// 'stop' pressed by client...example 2 sec delay...stop timestamp (time.Now()) called here
		// delay added 2 sec to total time so a false reading is given
		// must subtract delay from time.Now()

		for name, valueMap := range myMap {
			fmt.Println(name)
			for i := range valueMap {
				start := valueMap[i]["startTime"]
				stop := valueMap[i]["stopTime"]
				fmt.Println(start)
				// fmt.Println(stop  s.StopTime)
				fmt.Println(stop)
				// fmt.Println(s.StopTime)
				// fmt.Println(stop.Sub(s.StopTime))
				fmt.Println(stop.Sub(start)) // difference in js stop timestamp and time.Now() stop timesnamp
				fmt.Println(" ")
			}
		}

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
	http.ListenAndServe(":8081", nil) // launch server listening on port 8080
}

func main() {
	handleRequests()
}
