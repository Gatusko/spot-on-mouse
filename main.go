package main

import (
	"log"
	"runtime"
	"syscall"
	"unsafe"
)

// Win32 entry points. Resolved lazily on first call, so nothing happens at init.
//
// Every Win32 function has a doc page at a predictable URL:
//
//	learn.microsoft.com/en-us/windows/win32/api/<header>/nf-<header>-<funcname>
//
// Structs use ns- instead of nf-. Each page lists its header and the constants
// that go with it under "Requirements".
var (
	user32 = syscall.NewLazyDLL("user32.dll")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
	procRegisterClassEx = user32.NewProc("RegisterClassExW")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
	procCreateWindowEx = user32.NewProc("CreateWindowExW")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defwindowprocw
	procDefWindowProc = user32.NewProc("DefWindowProcW")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessage
	procGetMessage = user32.NewProc("GetMessageW")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessage
	procDispatchMessage = user32.NewProc("DispatchMessageW")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
	procPostQuitMessage = user32.NewProc("PostQuitMessage")
)

const (
	className  = "SpotOnMouseWindow"
	windowName = "Spot on Mouse"

	// WS_* window styles are a bitmask, combined with |.
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/window-styles
	//
	// WS_OVERLAPPEDWINDOW is itself six flags OR'd together:
	// WS_CAPTION|WS_SYSMENU|WS_THICKFRAME|WS_MINIMIZEBOX|WS_MAXIMIZEBOX.
	wsOverlappedWindow = 0x00CF0000
	wsVisible          = 0x10000000

	// CW_USEDEFAULT is a sentinel, not flags: (int)0x80000000 is INT_MIN, a
	// value no real screen coordinate can be. Only honoured for overlapped
	// windows; a WS_POPUP window treats it as zero.
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
	cwUseDefault = 0x80000000

	// CS_* class styles are a separate bitmask, used in WNDCLASSEXW.style.
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/window-class-styles
	csVRedraw = 0x0001
	csHRedraw = 0x0002

	// WM_* messages are ordinals, compared with == (never masked with &).
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-destroy
	wmDestroy = 0x0002

	wmCommand = 0x0111

	wsChild = 0x40000000

	// COLOR_* indexes the system colour table. WNDCLASSEXW.hbrBackground takes
	// either a real HBRUSH or one of these +1, since 0 already means "no brush".
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
	colorWindow = 5
	// BUTTON is a predefined class: Windows registered it at boot, so there is
	// no RegisterClassExW and no wndproc of our own for it.
	// https://learn.microsoft.com/en-us/windows/win32/controls/button-types-and-styles
	buttonClass = "BUTTON"

	// Our own control ID. Any value we pick; it comes back in WM_COMMAND.
	idToggleButton = 1001
)

// POINT.
// https://learn.microsoft.com/en-us/windows/win32/api/windef/ns-windef-point
type point struct {
	X int32
	Y int32
}

// MSG. The 4 bytes of padding after message are inserted by Go to match C.
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msg
type msg struct {
	hwnd    uintptr
	message uint32
	wParam  uintptr
	lParam  uintptr
	time    uint32
	pt      point
}

// WNDCLASSEXW. Field order and types must match the C struct exactly.
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-wndclassexw
type wndClassEx struct {
	size       uint32
	style      uint32
	wndProc    uintptr
	clsExtra   int32
	wndExtra   int32
	instance   uintptr
	icon       uintptr
	cursor     uintptr
	background uintptr
	menuName   *uint16
	className  *uint16
	iconSm     uintptr
}

// wndProc handles messages for our window and hands the rest back to Windows.
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-wndproc
func wndProc(hwnd uintptr, message uint32, wParam, lParam uintptr) uintptr {
	switch message {
	case wmDestroy:
		// Closing the window ends the message loop below.
		procPostQuitMessage.Call(0)
		return 0
	case wmCommand:
		// Low word of wParam is the control ID; high word is the notification
		// code (BN_CLICKED == 0 for buttons).
		if uint16(wParam) == idToggleButton {
			log.Println("button clicked")
		}
		return 0
	}

	ret, _, _ := procDefWindowProc.Call(hwnd, uintptr(message), wParam, lParam)
	return ret
}

func main() {
	// Window handles belong to the thread that created them, and so does the
	// message loop. Without this, Go may move us to another thread and the
	// window stops responding.
	// https://pkg.go.dev/runtime#LockOSThread
	runtime.LockOSThread()

	classNamePtr, err := syscall.UTF16PtrFromString(className)
	if err != nil {
		log.Fatalf("class name: %v", err)
	}
	windowNamePtr, err := syscall.UTF16PtrFromString(windowName)
	if err != nil {
		log.Fatalf("window name: %v", err)
	}

	wc := wndClassEx{
		style: csHRedraw | csVRedraw,
		// NewCallback turns a Go func into a C function pointer Windows can call.
		// https://pkg.go.dev/syscall#NewCallback
		wndProc:    syscall.NewCallback(wndProc),
		background: colorWindow + 1, // system color brush, so the window is visible
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
	atom, _, callErr := procRegisterClassEx.Call(uintptr(unsafe.Pointer(&wc)))
	if atom == 0 {
		log.Fatalf("RegisterClassExW: %v", callErr)
	}

	hwnd, _, callErr := procCreateWindowEx.Call(
		0, // no extended styles yet
		uintptr(unsafe.Pointer(classNamePtr)),
		uintptr(unsafe.Pointer(windowNamePtr)),
		wsOverlappedWindow|wsVisible,
		cwUseDefault, cwUseDefault, // position
		320, 240, // size
		0, // no parent
		0, // no menu
		0, // hInstance, unused for a single-exe app
		0, // no creation data
	)
	if hwnd == 0 {
		log.Fatalf("CreateWindowExW: %v", callErr)
	}

	buttonClassPtr, err := syscall.UTF16PtrFromString(buttonClass)
	if err != nil {
		log.Fatalf("button class: %v", err)
	}
	buttonTextPtr, err := syscall.UTF16PtrFromString("Enable spot")
	if err != nil {
		log.Fatalf("button text: %v", err)
	}

	button, _, callErr := procCreateWindowEx.Call(
		0,
		uintptr(unsafe.Pointer(buttonClassPtr)),
		uintptr(unsafe.Pointer(buttonTextPtr)),
		wsChild|wsVisible,
		20, 20, // position, relative to the parent's client area
		120, 32, // size
		hwnd,           // parent window
		idToggleButton, // hMenu slot doubles as the control ID for children
		0,
		0,
	)
	if button == 0 {
		log.Fatalf("CreateWindowExW(BUTTON): %v", callErr)
	}
	log.Println("window created, close it to exit")

	// The message pump. GetMessage blocks while the queue is empty, so this
	// loop idles at ~0% CPU.
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/using-messages-and-message-queues
	var m msg
	for {
		// GetMessage fills m in: Win32 returns data through out-parameters, so
		// we own the struct and hand Windows its address.
		ret, _, callErrInner := procGetMessage.Call(uintptr(unsafe.Pointer(&m)), 0, 0, 0)

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
		procDispatchMessage.Call(uintptr(unsafe.Pointer(&m)))
	}
}
