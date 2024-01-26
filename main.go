package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/something", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Something"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	log.Println("server is running on :80")

	log.Fatal(http.ListenAndServe(":80", nil))
}
