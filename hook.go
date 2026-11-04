package main

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/Gatusko/spot-on-mouse/constants"
	"github.com/Gatusko/spot-on-mouse/dll"
)

// MSLLHOOKSTRUCT — what lParam points at for a low-level mouse hook.
// 32 bytes on amd64; Go inserts the same padding before dwExtraInfo as C does.
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msllhookstruct
type msllHookStruct struct {
	pt          point
	mouseData   uint32
	flags       uint32
	time        uint32
	dwExtraInfo uintptr
}

var (
	mouseHook uintptr

	// Created ONCE, at package level. syscall.NewCallback allocates from a
	// small fixed pool that is never freed, so calling it on every install
	// would leak slots and eventually panic.
	mouseProcCallback = syscall.NewCallback(mouseProc)
)

// mouseProc runs for every mouse event on the desktop, on our own message
// thread. It must be quick and it must always pass the event along.
func mouseProc(nCode, wParam, lParam uintptr) uintptr {
	if int32(nCode) >= constants.HcAction && wParam == constants.WmMouseMove {
		info := (*msllHookStruct)(unsafe.Pointer(lParam))
		moveSpotTo(info.pt)
	}

	// Returning non-zero would swallow the event and break the mouse for the
	// whole system. Always hand on whatever the next hook says.
	ret, _, _ := dll.ProcCallNextHookEx.Call(0, nCode, wParam, lParam)
	return ret
}

// installMouseHook starts tracking. Returns false if the hook could not be set.
func installMouseHook() bool {
	hook, _, callErr := dll.ProcSetWindowsHookEx.Call(
		constants.WhMouseLL,
		mouseProcCallback,
		0, // hMod: no DLL to inject for a low-level hook — see if Windows agrees
		0, // 0 = every thread on the desktop, i.e. global
	)
	if hook == 0 {
		log.Printf("SetWindowsHookExW: %v", callErr)
		return false
	}

	mouseHook = hook
	return true
}

// removeMouseHook stops tracking.
func removeMouseHook() {
	if mouseHook == 0 {
		return
	}
	dll.ProcUnhookWindowsHookEx.Call(mouseHook)
	mouseHook = 0
}
