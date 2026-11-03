package constants

const (
	// BS_* button styles. BS_AUTORADIOBUTTON handles the checkmark and the
	// mutual exclusion for us; we only react to the resulting WM_COMMAND.
	// https://learn.microsoft.com/en-us/windows/win32/controls/button-styles
	BsAutoRadioButton = 0x00000009

	// WS_GROUP marks the first control of a group. Radio buttons are mutually
	// exclusive within one group, and the arrow keys move inside it.
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/window-styles
	WsGroup = 0x00020000

	// BM_SETCHECK / BST_CHECKED — tick a button programmatically.
	// https://learn.microsoft.com/en-us/windows/win32/controls/bm-setcheck
	BmSetCheck = 0x00F1
	BstChecked = 0x0001

	IdShapeCircle = 1004
	IdShapeSquare = 1005
)
