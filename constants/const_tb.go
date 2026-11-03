package constants

const (
	// The trackbar's predefined class, provided by comctl32.
	// https://learn.microsoft.com/en-us/windows/win32/controls/trackbar-control-reference
	TrackbarClass = "msctls_trackbar32"

	// ICC_BAR_CLASSES registers toolbar, statusbar, trackbar and tooltip.
	// https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-initcommoncontrolsex
	IccBarClasses = 0x00000004

	TbsNoTicks  = 0x0010 // a tick per pixel would be a solid bar
	TbsTooltips = 0x0100 // floating tooltip showing the value while dragging

	TbmSetPageSize = WmUser + 21

	// Control messages start at WM_USER; TBM_* are offsets from it.
	// https://learn.microsoft.com/en-us/windows/win32/controls/bumper-trackbar-control-reference-messages
	WmUser      = 0x0400
	TbmGetPos   = WmUser + 0
	TbmSetPos   = WmUser + 5
	TbmSetRange = WmUser + 6

	IdSizeSlider = 1003
)
