package main

import "syscall"

// GDI — drawing objects: brushes, pens, regions.
var (
	gdi32 = syscall.NewLazyDLL("gdi32.dll")

	// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createsolidbrush
	procCreateSolidBrush = gdi32.NewProc("CreateSolidBrush")

	// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
	procDeleteObject = gdi32.NewProc("DeleteObject")
)
