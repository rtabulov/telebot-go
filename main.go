package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// static folder
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port

	log.Print("Listening on http://localhost", port, "\n")
	log.Fatal(http.ListenAndServe(port, nil))
}
