package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	log.Println("Listening on port 8080")

	router.HandleFunc("/ping", pong)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalln("Failed to start server")
	}
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
