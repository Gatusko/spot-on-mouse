package main

const (
	// WS_* window styles are a bitmask, combined with |.
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/window-styles
	//
	// WS_OVERLAPPEDWINDOW is itself six flags OR'd together:
	// WS_CAPTION|WS_SYSMENU|WS_THICKFRAME|WS_MINIMIZEBOX|WS_MAXIMIZEBOX.
	wsOverlappedWindow = 0x00CF0000
	wsVisible          = 0x10000000
	wsChild            = 0x40000000
	wsPopup            = 0x80000000
)
