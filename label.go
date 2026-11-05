package main

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/Gatusko/spot-on-mouse/constants"
	"github.com/Gatusko/spot-on-mouse/dll"
)

// labels holds every label HWND, so wndProc can give only them a white
// background and leave the other controls alone.
var labels []uintptr

// isLabel reports whether hwnd is one of our labels.
func isLabel(hwnd uintptr) bool {
	for _, l := range labels {
		if l == hwnd {
			return true
		}
	}
	return false
}

// newLabel creates a static text label at the left margin.
func newLabel(parent uintptr, text string, y int) uintptr {
	classPtr, err := syscall.UTF16PtrFromString(constants.StaticClass)
	if err != nil {
		log.Fatalf("static class: %v", err)
	}
	textPtr, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		log.Fatalf("label %q: %v", text, err)
	}

	hwnd, _, callErr := dll.ProcCreateWindowEx.Call(
		0,
		uintptr(unsafe.Pointer(classPtr)),
		uintptr(unsafe.Pointer(textPtr)),
		constants.WsChild|constants.WsVisible,
		scaled(constants.Margin), scaled(y),
		scaled(constants.LabelWidth), scaled(constants.LabelHeight),
		parent,
		0, // labels take no input, so no control ID
		0,
		0,
	)
	if hwnd == 0 {
		log.Fatalf("CreateWindowExW(label %q): %v", text, callErr)
	}
	labels = append(labels, hwnd)
	return hwnd
}
