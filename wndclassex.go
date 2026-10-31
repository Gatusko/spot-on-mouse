package main

import (
	"github.com/Gatusko/spot-on-mouse/constants"
	"github.com/Gatusko/spot-on-mouse/dll"
)

// WNDCLASSEXW. Field order and types must match the C struct exactly.
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-wndclassexw
type wndClassEx struct {
	size       uint32
	style      uint32
	wndProc    uintptr
	clsExtra   int32
	wndExtra   int32
	instance   uintptr
	icon       uintptr
	cursor     uintptr
	background uintptr
	menuName   *uint16
	className  *uint16
	iconSm     uintptr
}

// wndProc handles messages for our window and hands the rest back to Windows.
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-wndproc
func wndProc(hwnd uintptr, message uint32, wParam, lParam uintptr) uintptr {
	switch message {
	case constants.WmDestroy:
		// Closing the window ends the message loop below.
		dll.ProcPostQuitMessage.Call(0)
		return 0
	case constants.WmCommand:
		// Low word of wParam is the control ID; lParam is the control's HWND.
		switch uint16(wParam) {
		case constants.IdToggleButton:
			setSpotEnabled(!spotEnabled, lParam)
		case constants.IdColorButton:
			pickSpotColor(hwnd)
		}
		return 0
	}

	ret, _, _ := dll.ProcDefWindowProc.Call(hwnd, uintptr(message), wParam, lParam)
	return ret
}
