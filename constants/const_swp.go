package constants

const (
	// SWP_* flags for SetWindowPos; a bitmask. Each one means "ignore the
	// corresponding argument".
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
	SwpNoSize     = 0x0001
	SwpNoZOrder   = 0x0004
	SwpNoActivate = 0x0010
	SwpNoMove     = 0x0002

	// HWND_TOPMOST is (HWND)-1. Go will not convert a negative constant to
	// uintptr, so write the all-bits-set pattern.
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
	HwndTopmost = ^uintptr(0)
)
