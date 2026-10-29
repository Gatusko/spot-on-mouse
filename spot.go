package main

import (
	"log"
	"syscall"
	"unsafe"
)

// The circle window. wndProc has to reach it to show and hide it, so it is a
// package-level var rather than a local in main.
var spotHwnd uintptr

// createSpot registers the spot's own window class and creates its window.
func createSpot() {
	classPtr, err := syscall.UTF16PtrFromString(spotClassName)
	if err != nil {
		log.Fatalf("spot class name: %v", err)
	}

	// A real HBRUSH this time, not a system colour index, so no +1.
	brush, _, callErr := procCreateSolidBrush.Call(spotColor)
	if brush == 0 {
		log.Fatalf("CreateSolidBrush: %v", callErr)
	}

	wc := wndClassEx{
		wndProc:    syscall.NewCallback(spotProc),
		background: brush,
		className:  classPtr,
	}
	wc.size = uint32(unsafe.Sizeof(wc))

	atom, _, callErr := procRegisterClassEx.Call(uintptr(unsafe.Pointer(&wc)))
	if atom == 0 {
		log.Fatalf("RegisterClassExW(spot): %v", callErr)
	}

	spotHwnd, _, callErr = procCreateWindowEx.Call(
		wsExLayered|wsExTransparent|wsExTopmost|wsExNoActivate|wsExToolWindow,
		uintptr(unsafe.Pointer(classPtr)),
		0, // no title — it has no caption to put one in
		wsPopup|wsVisible,
		200, 200, // fixed screen position for now, so we can find it
		spotSize, spotSize,
		0, // no parent
		0, // no menu
		0, // hInstance
		0, // no creation data
	)
	if spotHwnd == 0 {
		log.Fatalf("CreateWindowExW(spot): %v", callErr)
	}
	// A layered window draws nothing until its attributes are set.
	ret, _, callErr := procSetLayeredWindowAttributes.Call(
		spotHwnd,
		0,         // crKey, unused in LWA_ALPHA mode
		spotAlpha, // 0..255
		lwaAlpha,
	)
	if ret == 0 {
		log.Fatalf("SetLayeredWindowAttributes: %v", callErr)
	}
}

// spotProc: the background brush does the painting, the timer does the moving.
func spotProc(hwnd uintptr, message uint32, wParam, lParam uintptr) uintptr {
	switch message {
	case wmTimer:
		moveSpotToCursor(hwnd)
		return 0
	}

	ret, _, _ := procDefWindowProc.Call(hwnd, uintptr(message), wParam, lParam)
	return ret
}

// moveSpotToCursor centres the spot window on the mouse pointer.
func moveSpotToCursor(hwnd uintptr) {
	// Another out-parameter: we own the POINT, Windows fills it in.
	var pt point
	ret, _, callErr := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		log.Printf("GetCursorPos: %v", callErr)
		return
	}

	procSetWindowPos.Call(
		hwnd,
		0, // hWndInsertAfter, ignored because of SWP_NOZORDER
		uintptr(pt.X-spotSize/2),
		uintptr(pt.Y-spotSize/2),
		0, 0, // width/height, ignored because of SWP_NOSIZE
		swpNoSize|swpNoZOrder|swpNoActivate,
	)
}
