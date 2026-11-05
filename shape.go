package main

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/Gatusko/spot-on-mouse/constants"
	"github.com/Gatusko/spot-on-mouse/dll"
)

// createShapeRadios builds the circle/square pair and ticks the default.
func createShapeRadios(parent uintptr) {
	newLabel(parent, "Shape", constants.ShapeLabelY)
	circle := newRadio(parent, "Circle", constants.IdShapeCircle, constants.CircleRadioY, constants.WsGroup)
	newRadio(parent, "Square", constants.IdShapeSquare, constants.SquareRadioY, 0)

	// BS_AUTORADIOBUTTON only manages the tick once the user clicks, so the
	// initial state has to be set by hand to match spotCircle.
	dll.ProcSendMessage.Call(circle, constants.BmSetCheck, constants.BstChecked, 0)
}

// newRadio creates one auto-radio button. Same predefined BUTTON class as the
// push buttons — only the style differs.
func newRadio(parent uintptr, label string, id uintptr, y int, extra uintptr) uintptr {
	classPtr, err := syscall.UTF16PtrFromString(constants.ButtonClass)
	if err != nil {
		log.Fatalf("button class: %v", err)
	}
	labelPtr, err := syscall.UTF16PtrFromString(label)
	if err != nil {
		log.Fatalf("radio label %q: %v", label, err)
	}

	hwnd, _, callErr := dll.ProcCreateWindowEx.Call(
		0,
		uintptr(unsafe.Pointer(classPtr)),
		uintptr(unsafe.Pointer(labelPtr)),
		constants.WsChild|constants.WsVisible|constants.BsAutoRadioButton|extra,
		scaled(constants.Margin), scaled(y),
		scaled(constants.RadioWidth), scaled(constants.RadioHeight),
		parent,
		id,
		0,
		0,
	)
	if hwnd == 0 {
		log.Fatalf("CreateWindowExW(radio %q): %v", label, callErr)
	}
	return hwnd
}
