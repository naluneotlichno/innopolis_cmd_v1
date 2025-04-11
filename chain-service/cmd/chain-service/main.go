package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/szaluzhanskaya/Innopolis/chain-service/internal/controller/http"
)

func main() {

	port := "8080" //TODO: create ENV variable

	// Registers a handler for the /ping route
	http.HandleFunc("/ping", v1.PingHandler)

	// Starts the HTTP server on the port:8080
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
