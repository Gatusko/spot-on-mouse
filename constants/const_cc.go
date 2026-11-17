package constants

const (
	// CC_* flags for ChooseColorW.
	// https://learn.microsoft.com/en-us/windows/win32/api/commdlg/nf-commdlg-choosecolorw
	CcRgbInit    = 0x00000001 // start on result, instead of black
	CcFullOpen   = 0x00000002 // open with the custom-colour pane already expanded
	CcEnableHook = 0x00000010 // call the hook field with the dialog's messages
)

const (
	// Control IDs of the Red/Green/Blue edit boxes inside the colour dialog,
	// from colordlg.h. Every way of picking a colour rewrites these.
	// https://learn.microsoft.com/en-us/windows/win32/api/commdlg/nc-commdlg-lpcchookproc
	ColorRed   = 706
	ColorGreen = 707
	ColorBlue  = 708
)

const (
	// EN_* notifications, sent by an edit control to its parent as the high
	// word of WM_COMMAND's wParam.
	// https://learn.microsoft.com/en-us/windows/win32/controls/en-change
	EnChange = 0x0300
)
