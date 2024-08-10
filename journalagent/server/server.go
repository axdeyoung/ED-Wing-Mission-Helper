package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/axdeyoung/ed-wing-mission-helper/journalagent/journalparse"
)

const serverPort = 31173

var (
	server *http.Server
)

// InitServer is expected to be run when the program starts up.
func InitServer() {
	initRouteResponses()
	startServer(serverPort)
}

func startServer(port int) {
	server = &http.Server{Addr: fmt.Sprint(":", port)}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Error starting server: ", err)
		}
	}()
}

func initRouteResponses() {
	// note: the assignment of handler functions must be done by routes top-down.
	http.HandleFunc("/", respondStatus)
	http.HandleFunc("/trade/dump", respondTradeDump)
	http.HandleFunc("/trade/update", respondTradeUpdate)
}

func respondStatus(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "journal agent is operational.") //TODO: include more status responses ie: is Elite Dangerous running? is it receiving incoming connections from the website?
}

func respondTradeDump(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /trade/dump request\n")
	io.WriteString(w, journalparse.DumpTradeJson())
}

func respondTradeUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /trade/update request\n")
	io.WriteString(w, journalparse.UpdateTradeJson())
}
