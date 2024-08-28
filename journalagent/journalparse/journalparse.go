package journalparse

import (
	"os"
	"path/filepath"
)

var Dir string

func InitJournalParse() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		Dir = ""
	} else {
		Dir = filepath.Join(homeDir, "Saved Games", "Frontier Developments", "Elite Dangerous")
	}
}

func UserDataJson() string {
	return `{"data": "This is where the user data will go"}`
}
