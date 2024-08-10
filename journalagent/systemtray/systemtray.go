package systemtray

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
)

func InitSystray() {
	fmt.Println("Starting system tray interface")
	systray.Run(systrayReady, systrayExit)
}

func systrayReady() {
	fmt.Println("System tray ready. Populating with options...")
	systray.SetTitle("Elite Dangerous Helper Webserver")
	systray.SetTooltip("ED Wing Mission Helper")

	mOpenWebsite := systray.AddMenuItem("ED Wing Mission Helper", "Open ED Wing Mission Helper in your default browser")
	mFindJournal := systray.AddMenuItem("Journal Directory", "Locate Elite Dangerous journal directory")
	mSetPort := systray.AddMenuItem("Set Port", "Set the port the agent webserver is being hosted on.")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Shut down the server and terminate")

	go func() {
		fmt.Println("System tray prepared. Awaiting input...")
		for {
			select {
			case <-mOpenWebsite.ClickedCh:
				break // open the website in the OS's default browser
			case <-mFindJournal.ClickedCh:
				break // open a system GUI for the user to find their Elite Dangerous journal directory
			case <-mSetPort.ClickedCh:
				break // open a GUI for the user to enter a port for the server.
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func systrayExit() {
	os.Exit(0)
}
