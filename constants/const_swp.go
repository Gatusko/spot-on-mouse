package constants

const (
	// SWP_* flags for SetWindowPos; a bitmask. Each one means "ignore the
	// corresponding argument".
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
	SwpNoSize     = 0x0001
	SwpNoZOrder   = 0x0004
	SwpNoActivate = 0x0010
)
