package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("web")) // handler for web directory files
	http.Handle("/", fs)                   // registers FileServer as handler for all requests
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil) // launch server listening on port 8080
}
