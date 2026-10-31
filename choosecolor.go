package main

// CHOOSECOLORW. Field order and types must match the C struct exactly; Go
// inserts the same padding C does, giving 72 bytes on amd64.
// https://learn.microsoft.com/en-us/windows/win32/api/commdlg/ns-commdlg-choosecolorw
type chooseColor struct {
	structSize   uint32
	owner        uintptr
	instance     uintptr
	result       uint32
	custColors   *uint32
	flags        uint32
	custData     uintptr
	hook         uintptr
	templateName *uint16
}
