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
		case constants.IdShapeCircle:
			setSpotShape(true)
		case constants.IdShapeSquare:
			setSpotShape(false)
		}
		return 0
	case constants.WmHScroll:
		// lParam is the sending trackbar. A real scrollbar sends 0 here, which
		// is how you tell the two apart.
		if lParam != 0 {
			onSliderChange(lParam)
		}
		return 0
	case constants.WmDpiChanged:
		// Deliberately ignore the suggested rect: the panel keeps the pixel
		// size it was laid out at on startup, on every monitor.
		return 0
	case constants.WmCtlColorStatic:
		// Radios and the trackbar ask too; leave them to DefWindowProc.
		if !isLabel(lParam) {
			break
		}
		// Labels get the window's own white instead of the default grey,
		// with the text drawn transparent. The brush is system-owned: never
		// DeleteObject it.
		dll.ProcSetBkMode.Call(wParam, constants.BkTransparent)
		brush, _, _ := dll.ProcGetSysColorBrush.Call(constants.ColorWindow)
		return brush
	}

	ret, _, _ := dll.ProcDefWindowProc.Call(hwnd, uintptr(message), wParam, lParam)
	return ret
}
