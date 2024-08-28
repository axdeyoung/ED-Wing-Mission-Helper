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
	http.HandleFunc("/getuserdata", respondUserData)
}

func setHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func respondStatus(w http.ResponseWriter, r *http.Request) {
	setHeaders(&w)
	io.WriteString(w, "journal agent is operational.") //TODO: include more status responses ie: is Elite Dangerous running? is it receiving incoming connections from the website?
}

func respondUserData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /getuserdata request\n")
	setHeaders(&w)
	io.WriteString(w, journalparse.UserDataJson())
}
