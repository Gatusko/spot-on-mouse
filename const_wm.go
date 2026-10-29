package main

const (
	// WM_* messages are ordinals, compared with == (never masked with &).
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-destroy
	wmDestroy = 0x0002

	wmCommand = 0x0111

	wmTimer = 0x0113
)
