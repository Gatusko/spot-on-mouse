package dll

import "syscall"

// Common controls: trackbar, progress bar, toolbar and friends.
var (
	comctl32 = syscall.NewLazyDLL("comctl32.dll")

	// https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-initcommoncontrolsex
	ProcInitCommonControlsEx = comctl32.NewProc("InitCommonControlsEx")
)
