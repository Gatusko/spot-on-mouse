package constants

const (
	SpotClassName = "SpotOnMouseSpot"

	// The slider works directly in pixels: one unit per pixel, so it moves
	// smoothly instead of snapping to ten steps.
	SpotMinSize     = 8
	SpotMaxSize     = 120
	SpotDefaultSize = 36

	// COLORREF is 0x00BBGGRR — byte order is reversed from the RGB you'd write
	// in CSS. This is pure red.
	// https://learn.microsoft.com/en-us/windows/win32/gdi/colorref
	SpotDefaultColor = 0x000000FF

	// 0 = invisible, 255 = opaque.
	SpotAlpha = 160

	// LWA_ALPHA: use the alpha argument. (LWA_COLORKEY, 0x01, is the other mode.)
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setlayeredwindowattributes
	LwaAlpha = 0x00000002
)
