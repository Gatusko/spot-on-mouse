package main

const (
	// SWP_* flags for SetWindowPos; a bitmask. Each one means "ignore the
	// corresponding argument".
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
	swpNoSize     = 0x0001
	swpNoZOrder   = 0x0004
	swpNoActivate = 0x0010
)
