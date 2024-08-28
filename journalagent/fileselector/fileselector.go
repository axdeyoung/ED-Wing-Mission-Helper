package fileselector

import (
	"github.com/sqweek/dialog"
)

func LocateFolder(title string, previouslySelectedDir string) string {
	dir, err := dialog.Directory().Title(title).SetStartDir(previouslySelectedDir).Browse()
	if err != nil {
		return previouslySelectedDir
	}
	return dir
}
