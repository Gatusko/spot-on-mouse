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
)
