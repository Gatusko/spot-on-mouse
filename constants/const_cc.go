package constants

const (
	// CC_* flags for ChooseColorW.
	// https://learn.microsoft.com/en-us/windows/win32/api/commdlg/nf-commdlg-choosecolorw
	CcRgbInit  = 0x00000001 // start on result, instead of black
	CcFullOpen = 0x00000002 // open with the custom-colour pane already expanded
)
