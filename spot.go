package main

import (
	"log"
	"syscall"
	"unsafe"
)

// The circle window. wndProc has to reach it to show and hide it, so it is a
// package-level var rather than a local in main.

var (
	spotHwnd    uintptr
	spotEnabled bool

	// Current colour, and the brush painting it. The brush belongs to the
	// window class, so changing colour means swapping it and freeing the old one.
	spotColor = uint32(spotDefaultColor)
	spotBrush uintptr

	// ChooseColorW needs a 16-entry array for the user's custom colours and
	// FAILS if the pointer is nil. Package-level so picks survive between opens.
	customColors [16]uint32
)

// createSpot registers the spot's own window class and creates its window.
func createSpot() {
	classPtr, err := syscall.UTF16PtrFromString(spotClassName)
	if err != nil {
		log.Fatalf("spot class name: %v", err)
	}

	// A real HBRUSH this time, not a system colour index, so no +1.
	var callErr error
	spotBrush, _, callErr = procCreateSolidBrush.Call(uintptr(spotColor))
	if spotBrush == 0 {
		log.Fatalf("CreateSolidBrush: %v", callErr)
	}

	wc := wndClassEx{
		wndProc:    syscall.NewCallback(spotProc),
		background: spotBrush,
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
		0,        // no title — it has no caption to put one in
		wsPopup,  // start hidden
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

// setSpotEnabled shows or hides the spot and starts or stops its timer.
// buttonHwnd is the control whose label reflects the new state.
func setSpotEnabled(on bool, buttonHwnd uintptr) {
	spotEnabled = on

	if on {
		timer, _, callErr := procSetTimer.Call(spotHwnd, idSpotTimer, spotIntervalMs, 0)
		if timer == 0 {
			log.Printf("SetTimer: %v", callErr)
			return
		}

		// Jump to the cursor before showing, or the spot flashes at wherever
		// it was left last time.
		moveSpotToCursor(spotHwnd)
		procShowWindow.Call(spotHwnd, swShowNoActivate)
		setWindowText(buttonHwnd, "Disable spot")
		return
	}

	procKillTimer.Call(spotHwnd, idSpotTimer)
	procShowWindow.Call(spotHwnd, swHide)
	setWindowText(buttonHwnd, "Enable spot")
}

// setWindowText changes a window's title, which for a button is its label.
func setWindowText(hwnd uintptr, text string) {
	textPtr, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		log.Printf("window text %q: %v", text, err)
		return
	}
	procSetWindowText.Call(hwnd, uintptr(unsafe.Pointer(textPtr)))
}

// setSpotColor swaps the class background brush and repaints the spot.
func setSpotColor(color uint32) {
	newBrush, _, callErr := procCreateSolidBrush.Call(uintptr(color))
	if newBrush == 0 {
		log.Printf("CreateSolidBrush: %v", callErr)
		return
	}

	// The brush is a property of the window class, not the window.
	procSetClassLongPtr.Call(spotHwnd, gclpHbrBackground, newBrush)

	// Free the one we displaced, or we leak a GDI object on every pick.
	if spotBrush != 0 {
		procDeleteObject.Call(spotBrush)
	}
	spotBrush = newBrush
	spotColor = color

	// Nothing repaints on its own; ask for an erase and a paint.
	procInvalidateRect.Call(spotHwnd, 0, 1)
}

// pickSpotColor opens the standard Windows colour dialog.
func pickSpotColor(owner uintptr) {
	cc := chooseColor{
		owner:      owner,
		result:     spotColor,
		custColors: &customColors[0],
		flags:      ccRgbInit | ccFullOpen,
	}
	cc.structSize = uint32(unsafe.Sizeof(cc))

	ret, _, _ := procChooseColor.Call(uintptr(unsafe.Pointer(&cc)))
	if ret == 0 {
		// Also what Cancel returns, which is not an error worth logging.
		return
	}

	setSpotColor(cc.result)
}
