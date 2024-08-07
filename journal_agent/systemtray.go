package main

import (
	"github.com/getlantern/systray"
	"os"
)

func initSystray() {
	systray.Run(systrayReady, systrayExit)
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