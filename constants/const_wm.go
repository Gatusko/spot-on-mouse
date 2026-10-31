package constants

const (
	// WM_* messages are ordinals, compared with == (never masked with &).
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-destroy
	WmDestroy = 0x0002

	WmCommand = 0x0111

	WmTimer = 0x0113
)
