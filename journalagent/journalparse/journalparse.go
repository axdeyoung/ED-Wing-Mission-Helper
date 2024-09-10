package journalparse

import (
	"os"
	"path/filepath"
)

type genericEvent struct {
	EventType string `json:"event"`
}

var dir string

func InitJournalParse() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		dir = ""
	} else {
		UpdateDir(filepath.Join(homeDir, "Saved Games", "Frontier Developments", "Elite Dangerous"))
	}
}

func UpdateDir(newDir string) {
	dir = newDir
	// TODO: try to reopen all files in new directory.
	// TODO: notify user if any files fail to open. This may be the result of the game never having been played, an incorrect directory, or a non-existent directory.
}

func GetDir() string {
	return dir
}

func journalFile(fileName string) string {
	return filepath.Join(dir, fileName)
}

func UserDataJson() string {
	return `{"data": "This is where the user data will go"}`
}
