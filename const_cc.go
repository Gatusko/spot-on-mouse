package main

const (
	// CC_* flags for ChooseColorW.
	// https://learn.microsoft.com/en-us/windows/win32/api/commdlg/nf-commdlg-choosecolorw
	ccRgbInit  = 0x00000001 // start on result, instead of black
	ccFullOpen = 0x00000002 // open with the custom-colour pane already expanded
)
