package dll

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
	ProcRegisterClassEx = user32.NewProc("RegisterClassExW")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
	ProcCreateWindowEx = user32.NewProc("CreateWindowExW")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-defwindowprocw
	ProcDefWindowProc = user32.NewProc("DefWindowProcW")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessage
	ProcGetMessage = user32.NewProc("GetMessageW")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessage
	ProcDispatchMessage = user32.NewProc("DispatchMessageW")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
	ProcPostQuitMessage = user32.NewProc("PostQuitMessage")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorpos
	ProcGetCursorPos = user32.NewProc("GetCursorPos")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
	ProcSetWindowPos = user32.NewProc("SetWindowPos")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-settimer
	ProcSetTimer = user32.NewProc("SetTimer")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setlayeredwindowattributes
	ProcSetLayeredWindowAttributes = user32.NewProc("SetLayeredWindowAttributes")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
	ProcShowWindow = user32.NewProc("ShowWindow")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-killtimer
	ProcKillTimer = user32.NewProc("KillTimer")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowtextw
	ProcSetWindowText = user32.NewProc("SetWindowTextW")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setclasslongptrw
	ProcSetClassLongPtr = user32.NewProc("SetClassLongPtrW")

	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invalidaterect
	ProcInvalidateRect = user32.NewProc("InvalidateRect")
)
