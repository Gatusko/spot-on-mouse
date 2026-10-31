package main

import (
	"log"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/Gatusko/spot-on-mouse/constants"
	"github.com/Gatusko/spot-on-mouse/dll"
)

func main() {
	// Window handles belong to the thread that created them, and so does the
	// message loop. Without this, Go may move us to another thread and the
	// window stops responding.
	// https://pkg.go.dev/runtime#LockOSThread
	runtime.LockOSThread()

	classNamePtr, err := syscall.UTF16PtrFromString(constants.ClassName)
	if err != nil {
		log.Fatalf("class name: %v", err)
	}
	windowNamePtr, err := syscall.UTF16PtrFromString(constants.WindowName)
	if err != nil {
		log.Fatalf("window name: %v", err)
	}

	wc := wndClassEx{
		style: constants.CsHRedraw | constants.CsVRedraw,
		// NewCallback turns a Go func into a C function pointer Windows can call.
		// https://pkg.go.dev/syscall#NewCallback
		wndProc:    syscall.NewCallback(wndProc),
		background: constants.ColorWindow + 1, // system color brush, so the window is visible
		className:  classNamePtr,
	}
	wc.size = uint32(unsafe.Sizeof(wc))

	// uintptr(unsafe.Pointer(x)) must be written inline, inside the call. Hoisted
	// into a variable it becomes a plain integer the GC cannot see, and the thing
	// it points at may be moved or collected. See rule (4) of:
	// https://pkg.go.dev/unsafe#Pointer
	//
	// Win32 also reports failure through the return value; the error is only
	// meaningful once we know the call failed.
	atom, _, callErr := dll.ProcRegisterClassEx.Call(uintptr(unsafe.Pointer(&wc)))
	if atom == 0 {
		log.Fatalf("RegisterClassExW: %v", callErr)
	}

	hwnd, _, callErr := dll.ProcCreateWindowEx.Call(
		0, // no extended styles yet
		uintptr(unsafe.Pointer(classNamePtr)),
		uintptr(unsafe.Pointer(windowNamePtr)),
		constants.WsOverlappedWindow|constants.WsVisible,
		constants.CwUseDefault, constants.CwUseDefault, // position
		320, 240, // size
		0, // no parent
		0, // no menu
		0, // hInstance, unused for a single-exe app
		0, // no creation data
	)
	if hwnd == 0 {
		log.Fatalf("CreateWindowExW: %v", callErr)
	}

	buttonClassPtr, err := syscall.UTF16PtrFromString(constants.ButtonClass)
	if err != nil {
		log.Fatalf("button class: %v", err)
	}
	buttonTextPtr, err := syscall.UTF16PtrFromString("Enable spot")
	if err != nil {
		log.Fatalf("button text: %v", err)
	}

	button, _, callErr := dll.ProcCreateWindowEx.Call(
		0,
		uintptr(unsafe.Pointer(buttonClassPtr)),
		uintptr(unsafe.Pointer(buttonTextPtr)),
		constants.WsChild|constants.WsVisible,
		20, 20, // position, relative to the parent's client area
		120, 32, // size
		hwnd,                     // parent window
		constants.IdToggleButton, // hMenu slot doubles as the control ID for children
		0,
		0,
	)
	if button == 0 {
		log.Fatalf("CreateWindowExW(BUTTON): %v", callErr)
	}
	createSpot()

	colorTextPtr, err := syscall.UTF16PtrFromString("Choose colour...")
	if err != nil {
		log.Fatalf("colour button text: %v", err)
	}

	colorButton, _, callErr := dll.ProcCreateWindowEx.Call(
		0,
		uintptr(unsafe.Pointer(buttonClassPtr)), // same predefined class
		uintptr(unsafe.Pointer(colorTextPtr)),
		constants.WsChild|constants.WsVisible,
		20, 60, // below the toggle
		120, 32,
		hwnd,
		constants.IdColorButton,
		0,
		0,
	)
	if colorButton == 0 {
		log.Fatalf("CreateWindowExW(colour button): %v", callErr)
	}
	log.Println("window created, close it to exit")

	// The message pump. GetMessage blocks while the queue is empty, so this
	// loop idles at ~0% CPU.
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/using-messages-and-message-queues
	var m msg
	for {
		// GetMessage fills m in: Win32 returns data through out-parameters, so
		// we own the struct and hand Windows its address.
		ret, _, callErrInner := dll.ProcGetMessage.Call(uintptr(unsafe.Pointer(&m)), 0, 0, 0)

		// BOOL is a signed int32, but Call gives back an unsigned 64-bit
		// uintptr, so -1 arrives as 0x00000000FFFFFFFF. Truncating to int32
		// reinterprets it as the -1 Windows meant.
		if int32(ret) == -1 {
			log.Fatalf("GetMessageW: %v", callErrInner)
		}
		if ret == 0 { // WM_QUIT
			break
		}

		// This is what actually hands the message to wndProc above.
		dll.ProcDispatchMessage.Call(uintptr(unsafe.Pointer(&m)))
	}
}
