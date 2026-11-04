package main

import "github.com/Gatusko/spot-on-mouse/dll"

// uiDpi is the DPI of the monitor hosting the control panel; 96 means 100%.
var uiDpi = 96

// scaled converts a 96-DPI design value into device pixels. It returns uintptr
// because every caller is passing it straight into a Win32 call.
func scaled(px int) uintptr {
	return uintptr(px * uiDpi / 96)
}

// initUIScale reads the DPI of the monitor the window opened on. Call it after
// the window exists but before any child control is created.
func initUIScale(hwnd uintptr) {
	dpi, _, _ := dll.ProcGetDpiForWindow.Call(hwnd)

	// Zero means the call failed; the range check rejects nonsense.
	if dpi >= 48 && dpi <= 480 {
		uiDpi = int(dpi)
	}
}
