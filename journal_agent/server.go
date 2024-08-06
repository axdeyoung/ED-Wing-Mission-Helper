package main

import (
	"errors"
	"fmt"
	"os"
	"io"
	"net/http"
	"github.com/getlantern/systray"

	// local
	"journalparser"
)

func main() {
	http.HandleFunc("/trade/dump", getTradeDump)
	http.HandleFunc("/trade/update", getTradeUpdate)

	systray.Run(systrayReady, systrayExit)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

	
}

func systrayReady() {
	systray.SetTitle("Elite Dangerous Helper Webserver")
	systray.SetTooltip("Placeholder Tooltip")
	mQuit := systray.AddMenuItem("Quit", "Shut down the server and terminate")

	go func() {
		for {
			select{
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func systrayExit() {
	os.Exit(0)
}

func getTradeDump(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /trade/dump request\n")
	io.WriteString(w, journalparser.DumpTradeJson())
}

func getTradeUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /trade/update request\n")
	io.WriteString(w, journalparser.UpdateTradeJson())
}