package dll

import "syscall"

// Common dialogs: the standard Open / Save / Colour / Font boxes.
var (
	comdlg32 = syscall.NewLazyDLL("comdlg32.dll")

	// https://learn.microsoft.com/en-us/windows/win32/api/commdlg/nf-commdlg-choosecolorw
	ProcChooseColor = comdlg32.NewProc("ChooseColorW")
)
