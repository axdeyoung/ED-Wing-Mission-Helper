package journalparse

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/axdeyoung/ed-wing-mission-helper/journalagent/config"
)

type genericEvent struct {
	EventType string `json:"event"`
}

var dir string

var (
	journalReadSig      = make(chan struct{})
	journalQuitSig      = make(chan struct{})
	journalNewLinesChan = make(chan struct{})

	cargoReadSig      = make(chan struct{})
	cargoQuitSig      = make(chan struct{})
	cargoNewLinesChan = make(chan struct{})
)

func InitJournalParse() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		dir = ""
	} else {
		UpdateDir(filepath.Join(homeDir, config.Default_journal_dir_from_home))
	}
}

func UpdateDir(newDir string) {
	dir = newDir
	// TODO: try to reopen all files in new directory.
	// TODO: notify user if any files fail to open. This may be the result of the game never having been played, an incorrect directory, or a non-existent directory.
}

func newestJournalFilePath() (string, error) {
	var newestFilePath string
	var newestTime time.Time

	fileRegex, err := regexp.Compile(`^Journal\.(\d{4}-\d{2}-\d{2}T\d{2}\d{2}\d{2})\.01\.log$`)
	if err != nil {
		return "ERROR: FAILED TO COMPILE REGEX TO MATCH FILE", fmt.Errorf("failed to compile regex: %w", err)
	}

	err = filepath.WalkDir(dir, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// ignore directories
		if info.IsDir() {
			return nil
		}

		fileName := info.Name()

		matches := fileRegex.FindStringSubmatch(fileName)
		// ignore files that don't match the regex
		if len(matches) < 1 {
			return nil
		}

		// extract and parse date
		dateStampStr := matches[1]
		dateStamp, err := time.Parse("2006-01-02T150405", dateStampStr)
		if err != nil {
			return fmt.Errorf("failed to parse timestamp: %w", err)
		}

		// compare date and time and update newest time and file path
		if dateStamp.After(newestTime) {
			newestTime = dateStamp
			newestFilePath = path
		}

		return nil
	})
	if err != nil {
		return "", fmt.Errorf("error walking directory %s: %w", dir, err)
	}

	return newestFilePath, nil
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

func listenToFile(
	filePath string,
	isLog bool,
	readSigChan <-chan struct{},
	quitSigChan <-chan struct{},
	newLinesChan chan<- []string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	for {
		select {
		case <-readSigChan:
			// if this is not a log file, we need to read the whole file.
			// if this is a log file, we only want the new data since the last read.
			if !isLog {
				file.Seek(0, io.SeekStart)
			}
			scanner := bufio.NewScanner(file)
			var newLines []string
			// grab every line in the slice until the end of file.
			for scanner.Scan() {
				newLines = append(newLines, scanner.Text())
			}
			newLinesChan <- newLines
		case <-quitSigChan:
			return nil
		}
	}
}
