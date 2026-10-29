package main

// This is for OS windows
import "syscall"

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

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorpos
	procGetCursorPos = user32.NewProc("GetCursorPos")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
	procSetWindowPos = user32.NewProc("SetWindowPos")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-settimer
	procSetTimer = user32.NewProc("SetTimer")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setlayeredwindowattributes
	procSetLayeredWindowAttributes = user32.NewProc("SetLayeredWindowAttributes")
)
