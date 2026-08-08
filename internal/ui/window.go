//go:build windows && cgo

package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

const mainWindowTitle = "Music Basicifier"

func buildMainWindow(a fyne.App) fyne.Window {
	w := a.NewWindow(mainWindowTitle)
	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("https://www.youtube.com/watch?v=siCmqvfw_1g")
	audioFileEntry := widget.NewEntry()
	audioFileEntry.SetPlaceHolder("C:/source/music-basicifier/example-input/Useless-Station.wav")

	browseButton := widget.NewButton("Browse...", func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				return
			}
			audioFileEntry.SetText(reader.URI().Path())
			_ = reader.Close()
		}, w)
		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".wav", ".m4u", ".mp3"}))
		fileDialog.Show()
	})

	audioInput := container.NewBorder(nil, nil, nil, browseButton, audioFileEntry)
	confirmButton := widget.NewButton("Confirm", func() {
		dialog.ShowInformation("Input received", buildConfirmationMessage(urlEntry.Text, audioFileEntry.Text), w)
	})

	urlEntry.OnSubmitted = func(_ string) {
		confirmButton.OnTapped()
	}
	audioFileEntry.OnSubmitted = func(_ string) {
		confirmButton.OnTapped()
	}

	content := container.NewVBox(
		widget.NewLabel("Input"),
		widget.NewForm(
			widget.NewFormItem("YouTube URL", urlEntry),
			widget.NewFormItem("Audio File", audioInput),
		),
		confirmButton,
	)

	w.SetContent(content)
	return w
}
