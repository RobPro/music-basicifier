//go:build windows && cgo

package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const mainWindowTitle = "Music Basicifier"

func buildMainWindow(a fyne.App) fyne.Window {
	w := a.NewWindow(mainWindowTitle)
	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("https://www.youtube.com/watch?v=siCmqvfw_1g")

	content := container.NewVBox(
		widget.NewLabel("Input"),
		widget.NewForm(
			widget.NewFormItem("YouTube URL", urlEntry),
		),
	)

	w.SetContent(content)
	return w
}
