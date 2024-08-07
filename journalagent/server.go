package main

import (
	"fmt"
	"io"
	"net/http"
	"errors"
	"os"
)

const serverPortString = ":3333"
const serverAddress = nil // all available local addresses

func initServer() {
	initRouteResponses()

	err := http.ListenAndServe(serverPortString, serverAddress)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func initRouteResponses() {
	http.HandleFunc("/trade/dump", getTradeDump)
	http.HandleFunc("/trade/update", getTradeUpdate)
}

func getTradeDump(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /trade/dump request\n")
	io.WriteString(w, dumpTradeJson())
}

func getTradeUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /trade/update request\n")
	io.WriteString(w, updateTradeJson())
}