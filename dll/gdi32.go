package dll

import "syscall"

// GDI — drawing objects: brushes, pens, regions.
var (
	gdi32 = syscall.NewLazyDLL("gdi32.dll")

	// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createsolidbrush
	ProcCreateSolidBrush = gdi32.NewProc("CreateSolidBrush")

	// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
	ProcDeleteObject = gdi32.NewProc("DeleteObject")
)
