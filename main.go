package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

func main() {
	mux := http.NewServeMux()

	const filePathRoot = "."
	const port = "8080"

	handle := http.StripPrefix("/app", http.FileServer(http.Dir(filePathRoot)))
	apiCfg := apiConfig{}

	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handle))
	server := &http.Server{}
	server.Addr = ":" + port
	server.Handler = mux

	mux.HandleFunc("/healthz", handlerReadiness)
	mux.HandleFunc("/metrics", apiCfg.handlerCounter)
	mux.HandleFunc("/reset", apiCfg.handlerReset)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

type apiConfig struct {
	fileserverHits atomic.Int32
}
