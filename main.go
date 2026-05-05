package main

import (
	"log"
	"net/http"
)

func main() {
	servemux := http.NewServeMux()
	servemux.Handle("/", http.FileServer(http.Dir(".")))
	server := &http.Server{}
	server.Addr = ":8080"
	server.Handler = servemux

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
