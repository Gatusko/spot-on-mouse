package constants

const (
	// WM_* messages are ordinals, compared with == (never masked with &).
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-destroy
	WmDestroy = 0x0002

	WmCommand = 0x0111

	// A horizontal trackbar reports changes to its parent with WM_HSCROLL.
	// https://learn.microsoft.com/en-us/windows/win32/controls/wm-hscroll--trackbar-
	WmHScroll = 0x0114

	// https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-mousemove
	WmMouseMove = 0x0200

	// Sent when the window moves to a monitor with a different scale.
	// lParam points to a suggested RECT; DefWindowProc applies it.
	// https://learn.microsoft.com/en-us/windows/win32/hidpi/wm-dpichanged
	WmDpiChanged = 0x02E0

	// A static control (also trackbars and radio buttons) asks its parent for
	// a background brush. wParam is its HDC, lParam its HWND.
	// https://learn.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorstatic
	WmCtlColorStatic = 0x0138
)
