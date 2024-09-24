package journalparse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/axdeyoung/ed-wing-mission-helper/journalagent/config"
)

var dir string
var commanderName string

var (
	journalReadSig        = make(chan struct{})
	journalQuitSig        = make(chan struct{})
	journalNewLinesChan   = make(chan []string)
	journalListeningMutex sync.Mutex

	cargoReadSig        = make(chan struct{})
	cargoQuitSig        = make(chan struct{})
	cargoNewLinesChan   = make(chan []string)
	cargoListeningMutex sync.Mutex
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

	journalPath, err := newestJournalFilePath()
	if err != nil {
		fmt.Println("Unable to find newest Journal file path: ", err)
		// If there is no journal file for any reason, ignore the error
		// listenToFile will quietly fail and terminate. This is fine.
	}

	select {
	case journalQuitSig <- struct{}{}:
		// if there is journal file already open, close it
	default:
	}
	go listenToFile(journalPath, true, journalReadSig, journalQuitSig, journalNewLinesChan, &journalListeningMutex)

	select {
	case cargoQuitSig <- struct{}{}:
		// if there is cargo file already open, close it
	default:
	}
	go listenToFile(journalFile("Cargo.json"), false, cargoReadSig, cargoQuitSig, cargoNewLinesChan, &cargoListeningMutex)
	fmt.Println("Opened Journal Path: ", journalPath)
	// TODO: notify user if any files fail to open. This may be the result of the game never having been played, an incorrect directory, or a non-existent directory.
}

func newestJournalFilePath() (string, error) {
	var newestFilePath string
	var newestTime time.Time
	newestTime = time.Unix(0, 0)

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
		fmt.Println("Error finding journal file: ", err)
		return "", fmt.Errorf("error walking directory %s: %w", dir, err)
	}

	fmt.Println("Found journal path: ", newestFilePath)
	return newestFilePath, nil
}

func GetDir() string {
	return dir
}

func GetCommanderName() string {
	return commanderName
}
func journalFile(fileName string) string {
	return filepath.Join(dir, fileName)
}

func UserDataJson() string {
	// structure is so simple it should be fine to just marshal it manually.
	return fmt.Sprintf(`{ "Name":"%s" }`, commanderName)
}

func listenToFile(
	filePath string,
	isLog bool,
	readSigChan <-chan struct{},
	quitSigChan <-chan struct{},
	newLinesChan chan<- []string,
	fileMutex *sync.Mutex) error {

	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	fmt.Println("Opened file and started listening: ", filePath)
	defer fmt.Println("Stopped listening and closed file: ", filePath)
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

func GetUpdatesJson() (string, error) {
	cargoToChange := false
	updatesToSend := ""

	select {
	case journalReadSig <- struct{}{}:
	default:
		return "", fmt.Errorf("error reading journal file")
	}
	latestJournalEntriesSliceRaw := <-journalNewLinesChan

	for _, rawEntry := range latestJournalEntriesSliceRaw {
		var entry map[string]any
		err := json.Unmarshal([]byte(rawEntry), &entry)
		if err != nil {
			fmt.Println("Warning: error unmarshalling journal entry: ", rawEntry)
			continue
		}

		newUpdateToSend, cargoChanged, err := filterAndParseJournalEntry(entry)
		if err != nil {
			fmt.Println("Warning: error parsing journal entry: ", rawEntry, ": ", err)
			continue
		}
		if newUpdateToSend != "" {
			updatesToSend += newUpdateToSend + "\n"
		}
		if cargoChanged {
			cargoToChange = true
		}
	}
	if cargoToChange {
		newCargoToSend, err := parseCargo()
		if err != nil {
			return "", fmt.Errorf("unable to parse cargo: %w", err)
		}
		updatesToSend += newCargoToSend + "\n"
	}
	return updatesToSend, nil
}

/*
returns 3 values:

	string: a JSON entry that is to be appended to the data sent to the server. May be empty.
	bool: Does the given entry indicate a cargo update?
	error: any error causing this function to fail.
*/
func filterAndParseJournalEntry(entry map[string]any) (string, bool, error) {
	var err error
	var dataToSend []byte
	var updateCargo bool

	err = nil
	updateCargo = false

	switch entry["event"] {
	case "Commander":
		commanderName = entry["Name"].(string)
		delete(entry, "FID")
		dataToSend, err = json.Marshal(entry)
	case "Cargo":
		updateCargo = true
	default: // not an event type we care about. Ignore it.
		return "", false, nil
	}

	// aggregating all the errors here instead of having a separate one in each case.
	if err != nil {
		return "", false, fmt.Errorf("error processing journal entry: %w", err)
	}

	return string(dataToSend), updateCargo, err
}

func parseCargo() (string, error) {
	cargoReadSig <- struct{}{}
	cargoLines := <-cargoNewLinesChan
	cargoJson := strings.Join(cargoLines, "\n")
	return cargoJson, nil
}
