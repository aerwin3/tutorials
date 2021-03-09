package api

import (
	"log"
	"net/http"
)

func registerHandlers() {
	http.HandleFunc("/", indexHandler)
}

func Start() {
	registerHandlers()
	log.Fatal(http.ListenAndServe(":10000", nil))
}
