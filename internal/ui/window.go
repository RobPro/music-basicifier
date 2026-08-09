//go:build windows && cgo

package ui

import (
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/RobPro/music-basicifier/internal/platform"
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
	outputPreview := widget.NewMultiLineEntry()
	outputPreview.SetText(buildEmptyPreviewText())
	outputPreview.Disable()
	previewStatus := widget.NewLabel(buildPreviewStatusText(false))
	confirmButton := widget.NewButton("Confirm", func() {
		if url := urlEntry.Text; url != "" {
			downloadDir := filepath.Join("C:/", "source", "music-basicifier", "downloads")
			outputPath, err := platform.DownloadYouTubeAudio(url, downloadDir)
			if err != nil {
				wrapped := fmt.Errorf("could not download YouTube audio: %w", err)
				logErrorWithContext("download YouTube audio", err)
				dialog.ShowError(fmt.Errorf("%s", buildErrorMessage(wrapped)), w)
				return
			}
			dialog.ShowInformation("Input received", buildConfirmationMessage(url, outputPath), w)
			return
		}

		if err := platform.ValidateAudioFilePath(audioFileEntry.Text); err != nil {
			wrapped := fmt.Errorf("invalid audio file: %w", err)
			logErrorWithContext("validate audio file", err)
			dialog.ShowError(fmt.Errorf("%s", buildErrorMessage(wrapped)), w)
			return
		}

		input, err := buildConversionInputFromAudioFile(audioFileEntry.Text)
		if err != nil {
			wrapped := fmt.Errorf("could not extract melody: %w", err)
			logErrorWithContext("extract melody", err)
			dialog.ShowError(fmt.Errorf("%s", buildErrorMessage(wrapped)), w)
			return
		}
		bundle, err := platform.GenerateOutputBundle(input)
		if err != nil {
			wrapped := fmt.Errorf("could not generate outputs: %w", err)
			logErrorWithContext("generate outputs", err)
			dialog.ShowError(fmt.Errorf("%s", buildErrorMessage(wrapped)), w)
			return
		}
		outputPreview.SetText(buildPreviewText(bundle))
		previewStatus.SetText(buildPreviewStatusText(true))
		dialog.ShowInformation("Input received", buildConfirmationMessage(urlEntry.Text, audioFileEntry.Text), w)
	})

	urlEntry.OnSubmitted = func(_ string) {
		confirmButton.OnTapped()
	}
	audioFileEntry.OnSubmitted = func(_ string) {
		confirmButton.OnTapped()
	}

	copyButton := widget.NewButton("Copy Output", func() {
		copyText, err := buildCopyText(outputPreview.Text)
		if err != nil {
			wrapped := fmt.Errorf("could not copy output: %w", err)
			logErrorWithContext("copy output", err)
			dialog.ShowError(fmt.Errorf("%s", buildErrorMessage(wrapped)), w)
			return
		}
		if err := platform.CopyTextToClipboard(copyText); err != nil {
			wrapped := fmt.Errorf("could not copy output: %w", err)
			logErrorWithContext("copy output", err)
			dialog.ShowError(fmt.Errorf("%s", buildErrorMessage(wrapped)), w)
			return
		}
		dialog.ShowInformation("Copied", "Output copied to clipboard", w)
	})

	content := container.NewVBox(
		widget.NewLabel("Input"),
		widget.NewForm(
			widget.NewFormItem("YouTube URL", urlEntry),
			widget.NewFormItem("Audio File", audioInput),
		),
		confirmButton,
		widget.NewLabel("Output Preview"),
		outputPreview,
		previewStatus,
		copyButton,
	)

	w.SetContent(content)
	return w
}
