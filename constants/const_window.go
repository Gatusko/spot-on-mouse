package constants

const (
	ClassName  = "SpotOnMouseWindow"
	WindowName = "Spot on Mouse"

	// CW_USEDEFAULT is a sentinel, not flags: (int)0x80000000 is INT_MIN, a
	// value no real screen coordinate can be. Only honoured for overlapped
	// windows; a WS_POPUP window treats it as zero.
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
	CwUseDefault = 0x80000000

	// COLOR_* indexes the system colour table. WNDCLASSEXW.hbrBackground takes
	// either a real HBRUSH or one of these +1, since 0 already means "no brush".
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
	ColorWindow = 5
	// BUTTON is a predefined class: Windows registered it at boot, so there is
	// no RegisterClassExW and no wndproc of our own for it.
	// https://learn.microsoft.com/en-us/windows/win32/controls/button-types-and-styles
	ButtonClass = "BUTTON"

	// Our own control ID. Any value we pick; it comes back in WM_COMMAND.
	IdToggleButton = 1001

	IdColorButton = 1002

	// GCLP_HBRBACKGROUND is -10. Go refuses to convert a negative constant to
	// uintptr, so write the two's-complement bit pattern instead: ^9 == -10.
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setclasslongptrw
	GclpHbrBackground = ^uintptr(9)

	// DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2 is -4. Go will not convert a
	// negative constant to uintptr, so write the two's-complement pattern: ^3 == -4.
	// https://learn.microsoft.com/en-us/windows/win32/hidpi/dpi-awareness-context
	DpiPerMonitorAwareV2 = ^uintptr(3)

	// STATIC is the predefined class for text labels. Default style SS_LEFT
	// is 0, so WS_CHILD|WS_VISIBLE alone gives left-aligned text.
	// https://learn.microsoft.com/en-us/windows/win32/controls/static-controls
	StaticClass = "STATIC"

	// TRANSPARENT background mode: text is drawn without filling behind it.
	// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
	BkTransparent = 1
)
