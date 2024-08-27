package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Prompt(message string) string {
	a := app.New()
	window := a.NewWindow("NewWindow() argument")

	hello := widget.NewLabel("NewLabel() argument")
	window.SetContent(container.NewVBox(hello, widget.NewButton(
		"NewButton() argument", func() {
			hello.SetText("SetText() argument")
		}),
	))
	window.ShowAndRun()
	return "Prompt() placeholderString"
}
