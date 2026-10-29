package main

const (
	// WS_EX_* extended styles — yet another separate bitmask, passed as the
	// FIRST argument to CreateWindowExW (not the style argument).
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/extended-window-styles
	wsExTopmost     = 0x00000008
	wsExTransparent = 0x00000020
	wsExLayered     = 0x00080000
	wsExToolWindow  = 0x00000080
	wsExNoActivate  = 0x08000000
)
