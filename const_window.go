package main

const (
	className  = "SpotOnMouseWindow"
	windowName = "Spot on Mouse"

	// CW_USEDEFAULT is a sentinel, not flags: (int)0x80000000 is INT_MIN, a
	// value no real screen coordinate can be. Only honoured for overlapped
	// windows; a WS_POPUP window treats it as zero.
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
	cwUseDefault = 0x80000000

	// COLOR_* indexes the system colour table. WNDCLASSEXW.hbrBackground takes
	// either a real HBRUSH or one of these +1, since 0 already means "no brush".
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
	colorWindow = 5
	// BUTTON is a predefined class: Windows registered it at boot, so there is
	// no RegisterClassExW and no wndproc of our own for it.
	// https://learn.microsoft.com/en-us/windows/win32/controls/button-types-and-styles
	buttonClass = "BUTTON"

	// Our own control ID. Any value we pick; it comes back in WM_COMMAND.
	idToggleButton = 1001
)
