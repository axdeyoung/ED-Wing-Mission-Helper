package main

import (
	"github.com/axdeyoung/ed-wing-mission-helper/journalagent/gui"
	"github.com/axdeyoung/ed-wing-mission-helper/journalagent/server"
	"github.com/axdeyoung/ed-wing-mission-helper/journalagent/systemtray"
)

func main() {
	server.InitServer()
	systemtray.InitSystray()
	gui.Prompt("Test String")
}
