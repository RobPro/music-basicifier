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

func buildConversionInputFromAudioFile(path string) (*platform.ConversionInput, error) {
	melody, err := platform.ExtractMelodyFromAudioFile(path)
	if err != nil {
		return nil, err
	}
	return platform.BuildConversionInput(melody)
}

func buildPreviewText(bundle *platform.OutputBundle) string {
	if bundle == nil {
		return ""
	}
	return bundle.QBASIC + "\n\n" + bundle.Adafruit
}

func buildCopyText(preview string) string {
	return preview
}

func buildErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func buildEmptyPreviewText() string {
	return "No output yet. Confirm an input to generate preview text."
}

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
	confirmButton := widget.NewButton("Confirm", func() {
		if url := urlEntry.Text; url != "" {
			downloadDir := filepath.Join("C:/", "source", "music-basicifier", "downloads")
			outputPath, err := platform.DownloadYouTubeAudio(url, downloadDir)
			if err != nil {
				_ = platform.LogError(fmt.Sprintf("could not download YouTube audio: %v", err))
				dialog.ShowError(fmt.Errorf("could not download YouTube audio: %w", err), w)
				return
			}
			dialog.ShowInformation("Input received", buildConfirmationMessage(url, outputPath), w)
			return
		}

		if err := platform.ValidateAudioFilePath(audioFileEntry.Text); err != nil {
			_ = platform.LogError(fmt.Sprintf("invalid audio file: %v", err))
			dialog.ShowError(fmt.Errorf("invalid audio file: %w", err), w)
			return
		}

		input, err := buildConversionInputFromAudioFile(audioFileEntry.Text)
		if err != nil {
			_ = platform.LogError(fmt.Sprintf("could not extract melody: %v", err))
			dialog.ShowError(fmt.Errorf("could not extract melody: %s", buildErrorMessage(err)), w)
			return
		}
		bundle, err := platform.GenerateOutputBundle(input)
		if err != nil {
			_ = platform.LogError(fmt.Sprintf("could not generate outputs: %v", err))
			dialog.ShowError(fmt.Errorf("could not generate outputs: %s", buildErrorMessage(err)), w)
			return
		}
		outputPreview.SetText(buildPreviewText(bundle))
		dialog.ShowInformation("Input received", buildConfirmationMessage(urlEntry.Text, audioFileEntry.Text), w)
	})

	urlEntry.OnSubmitted = func(_ string) {
		confirmButton.OnTapped()
	}
	audioFileEntry.OnSubmitted = func(_ string) {
		confirmButton.OnTapped()
	}

	copyButton := widget.NewButton("Copy Output", func() {
		if err := platform.CopyTextToClipboard(buildCopyText(outputPreview.Text)); err != nil {
			dialog.ShowError(fmt.Errorf("could not copy output: %w", err), w)
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
		copyButton,
	)

	w.SetContent(content)
	return w
}
