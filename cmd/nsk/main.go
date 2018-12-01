package main

import (
	"log"
	"net/http"
	"fmt"
	"runtime"
)

func main() {
	log.Print("Starting application...")

	port := "8080"

	s := http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	http.HandleFunc("/", handler)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Serever is stoped with error: %v", err)
	}

    log.Print("Application has been stoped")
}


func handler  (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, runtime.Version())
}