package constants

const (
	// WH_MOUSE_LL — a global low-level mouse hook. Unlike most global hooks it
	// needs no DLL: there is no injection, and the callback runs on the thread
	// that installed it.
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/about-hooks
	WhMouseLL = 14

	// Codes below zero must be passed straight on without being examined.
	HcAction = 0
)
