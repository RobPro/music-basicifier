//go:build windows && cgo

// Package ui contains Fyne application setup and top-level window wiring.
package ui

import "fyne.io/fyne/v2/app"

// Run starts the desktop application event loop.
func Run() error {
	a := app.New()
	w := buildMainWindow(a)
	w.ShowAndRun()
	return nil
}
