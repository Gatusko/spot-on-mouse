package main

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/Gatusko/spot-on-mouse/constants"
	"github.com/Gatusko/spot-on-mouse/dll"
)

// INITCOMMONCONTROLSEX — two DWORDs, no padding.
// https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-initcommoncontrolsex
type initCommonControlsEx struct {
	size uint32
	icc  uint32
}

// initCommonControls registers the comctl32 window classes. Without it the
// trackbar class does not exist and CreateWindowExW fails.
func initCommonControls() {
	icc := initCommonControlsEx{icc: constants.IccBarClasses}
	icc.size = uint32(unsafe.Sizeof(icc))

	ret, _, callErr := dll.ProcInitCommonControlsEx.Call(uintptr(unsafe.Pointer(&icc)))
	if ret == 0 {
		log.Fatalf("InitCommonControlsEx: %v", callErr)
	}
}

// createSizeSlider builds the 1..10 trackbar that sets the spot diameter.
func createSizeSlider(parent uintptr) uintptr {
	newLabel(parent, "Size", constants.SizeLabelY)

	classPtr, err := syscall.UTF16PtrFromString(constants.TrackbarClass)
	if err != nil {
		log.Fatalf("trackbar class: %v", err)
	}

	slider, _, callErr := dll.ProcCreateWindowEx.Call(
		0,
		uintptr(unsafe.Pointer(classPtr)),
		0, // no text
		constants.WsChild|constants.WsVisible|constants.TbsNoTicks|constants.TbsTooltips,
		scaled(constants.Margin), scaled(constants.SliderY),
		scaled(constants.SliderWidth), scaled(constants.SliderHeight),
		parent,
		constants.IdSizeSlider,
		0,
		0,
	)
	if slider == 0 {
		log.Fatalf("CreateWindowExW(trackbar): %v", callErr)
	}

	// Range is in pixels now. TBM_SETRANGE still packs min into the low word
	// of lParam and max into the high word.
	dll.ProcSendMessage.Call(
		slider,
		constants.TbmSetRange,
		1, // redraw
		uintptr(constants.SpotMinSize)|uintptr(constants.SpotMaxSize)<<16,
	)
	dll.ProcSendMessage.Call(slider, constants.TbmSetPos, 1, constants.SpotDefaultSize)

	// Clicking the channel jumps by 8px rather than a fifth of the range.
	dll.ProcSendMessage.Call(slider, constants.TbmSetPageSize, 0, 8)

	return slider
}

func onSliderChange(slider uintptr) {
	pos, _, _ := dll.ProcSendMessage.Call(slider, constants.TbmGetPos, 0, 0)
	setSpotSizePx(int32(pos))
}
