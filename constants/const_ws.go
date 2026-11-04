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

	// The pieces WS_OVERLAPPEDWINDOW is made of, so styles can be composed
	// rather than copied as one magic number.
	WsCaption     = 0x00C00000
	WsSysMenu     = 0x00080000
	WsThickFrame  = 0x00040000 // the resize grip
	WsMinimizeBox = 0x00020000
	WsMaximizeBox = 0x00010000

	// A fixed-size panel: overlapped, minus the resize border and the
	// maximise box. The layout is hardcoded, so resizing has nothing to do.
	WsFixedWindow = WsCaption | WsSysMenu | WsMinimizeBox
)
