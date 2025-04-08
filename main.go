package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "ok")
		if err != nil {
			log.Fatal(err)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
