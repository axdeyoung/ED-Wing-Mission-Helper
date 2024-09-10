package systemtray

import (
	"fmt"
	"os"

	"fyne.io/systray"
	"github.com/axdeyoung/ed-wing-mission-helper/journalagent/fileselector"
	"github.com/axdeyoung/ed-wing-mission-helper/journalagent/journalparse"
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
				fmt.Println("Previous journal directory: ", journalparse.GetDir())
				journalparse.UpdateDir(fileselector.LocateFolder("Locate Journal Directory", journalparse.GetDir())) // open a system GUI for the user to find their Elite Dangerous journal directory
				fmt.Println("New journal directory: ", journalparse.GetDir())
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
