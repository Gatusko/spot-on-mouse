package constants

const (
	SpotClassName = "SpotOnMouseSpot"

	// Diameter in pixels.
	SpotSize = 40

	// COLORREF is 0x00BBGGRR — byte order is reversed from the RGB you'd write
	// in CSS. This is pure red.
	// https://learn.microsoft.com/en-us/windows/win32/gdi/colorref
	SpotDefaultColor = 0x000000FF

	// Timer ID. Ours to choose; only has to be unique per window.
	IdSpotTimer = 1

	// Milliseconds between position updates — roughly 60 per second.
	SpotIntervalMs = 16

	// 0 = invisible, 255 = opaque.
	SpotAlpha = 160

	// LWA_ALPHA: use the alpha argument. (LWA_COLORKEY, 0x01, is the other mode.)
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setlayeredwindowattributes
	LwaAlpha = 0x00000002
)
