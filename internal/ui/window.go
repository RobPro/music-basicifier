//go:build windows && cgo

package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

const mainWindowTitle = "Music Basicifier"

func buildMainWindow(a fyne.App) fyne.Window {
	w := a.NewWindow(mainWindowTitle)
	w.SetContent(container.NewWithoutLayout())
	return w
}
