package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/axdeyoung/ed-wing-mission-helper/journalagent/journalparse"
)

const serverPortString = ":3333"

func InitServer() {
	initRouteResponses()

	err := http.ListenAndServe(serverPortString, nil)
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
	io.WriteString(w, journalparse.DumpTradeJson())
}

func getTradeUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /trade/update request\n")
	io.WriteString(w, journalparse.UpdateTradeJson())
}
