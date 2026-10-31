package constants

const (
	// WS_EX_* extended styles — yet another separate bitmask, passed as the
	// FIRST argument to CreateWindowExW (not the style argument).
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/extended-window-styles
	WsExTopmost     = 0x00000008
	WsExTransparent = 0x00000020
	WsExLayered     = 0x00080000
	WsExToolWindow  = 0x00000080
	WsExNoActivate  = 0x08000000
)
