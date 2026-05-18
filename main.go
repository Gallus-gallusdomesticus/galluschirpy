package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	server := &http.Server{}
	server.Addr = ":8080"
	server.Handler = mux

	mux.HandleFunc("/healthz", handlerReadiness)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
