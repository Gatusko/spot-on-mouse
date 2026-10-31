package constants

const (
	// WS_* window styles are a bitmask, combined with |.
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/window-styles
	//
	// WS_OVERLAPPEDWINDOW is itself six flags OR'd together:
	// WS_CAPTION|WS_SYSMENU|WS_THICKFRAME|WS_MINIMIZEBOX|WS_MAXIMIZEBOX.
	WsOverlappedWindow = 0x00CF0000
	WsVisible          = 0x10000000
	WsChild            = 0x40000000
	WsPopup            = 0x80000000
)
