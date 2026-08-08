//go:build !cgo

// Package ui contains Fyne application setup and top-level window wiring.
package ui

import "errors"

// Run returns an error when the binary is built without CGO support.
func Run() error {
	return errors.New("ui requires cgo-enabled build on Windows")
}
