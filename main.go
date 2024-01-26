package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/something", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Something"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}
	httpPort := fmt.Sprintf(":%s", Port)
	log.Println("server is running on ", httpPort)

	log.Fatal(http.ListenAndServe(httpPort, nil))
}
