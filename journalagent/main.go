package main

import (
	"github.com/axdeyoung/ed-wing-mission-helper/journalagent/journalparse"
	"github.com/axdeyoung/ed-wing-mission-helper/journalagent/server"
	"github.com/axdeyoung/ed-wing-mission-helper/journalagent/systemtray"
)

func main() {
	journalparse.InitJournalParse()
	server.InitServer()
	systemtray.InitSystray()
}
