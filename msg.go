package main

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
