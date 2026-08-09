package platform

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	cfUnicodeText = 13
	gmemMoveable  = 0x0002
)

var (
	user32               = windows.NewLazySystemDLL("user32.dll")
	kernel32             = windows.NewLazySystemDLL("kernel32.dll")
	procOpenClipboard    = user32.NewProc("OpenClipboard")
	procCloseClipboard   = user32.NewProc("CloseClipboard")
	procEmptyClipboard   = user32.NewProc("EmptyClipboard")
	procSetClipboardData = user32.NewProc("SetClipboardData")
	procGlobalAlloc      = kernel32.NewProc("GlobalAlloc")
	procGlobalLock       = kernel32.NewProc("GlobalLock")
	procGlobalUnlock     = kernel32.NewProc("GlobalUnlock")
	procGlobalFree       = kernel32.NewProc("GlobalFree")
	procRtlMoveMemory    = kernel32.NewProc("RtlMoveMemory")
)

// CopyTextToClipboard copies a string onto the Windows clipboard using Win32 APIs.
func CopyTextToClipboard(text string) error {
	if err := validateClipboardText(text); err != nil {
		return err
	}

	if r1, _, err := procOpenClipboard.Call(0); r1 == 0 {
		return fmt.Errorf("open clipboard failed: %w", err)
	}
	defer func() {
		_, _, _ = procCloseClipboard.Call()
	}()

	if r1, _, err := procEmptyClipboard.Call(); r1 == 0 {
		return fmt.Errorf("clear clipboard failed: %w", err)
	}

	utf16, err := windows.UTF16FromString(text)
	if err != nil {
		return fmt.Errorf("encode clipboard text failed: %w", err)
	}
	byteLen := uintptr(len(utf16) * 2)
	hMem, _, err := procGlobalAlloc.Call(gmemMoveable, byteLen)
	if hMem == 0 {
		return fmt.Errorf("allocate clipboard memory failed: %w", err)
	}

	locked, _, err := procGlobalLock.Call(hMem)
	if locked == 0 {
		_, _, _ = procGlobalFree.Call(hMem)
		return fmt.Errorf("lock clipboard memory failed: %w", err)
	}

	_, _, _ = procRtlMoveMemory.Call(locked, uintptr(unsafe.Pointer(&utf16[0])), byteLen)
	if r1, _, err := procGlobalUnlock.Call(hMem); r1 == 0 && err != windows.ERROR_SUCCESS {
		_, _, _ = procGlobalFree.Call(hMem)
		return fmt.Errorf("unlock clipboard memory failed: %w", err)
	}

	if r1, _, err := procSetClipboardData.Call(cfUnicodeText, hMem); r1 == 0 {
		_, _, _ = procGlobalFree.Call(hMem)
		return fmt.Errorf("set clipboard data failed: %w", err)
	}

	return nil
}

func validateClipboardText(text string) error {
	if text == "" {
		return fmt.Errorf("text is required")
	}
	return nil
}
