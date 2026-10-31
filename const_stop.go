package main

const (
	spotClassName = "SpotOnMouseSpot"

	// Diameter in pixels.
	spotSize = 40

	// COLORREF is 0x00BBGGRR — byte order is reversed from the RGB you'd write
	// in CSS. This is pure red.
	// https://learn.microsoft.com/en-us/windows/win32/gdi/colorref
	spotDefaultColor = 0x000000FF

	// Timer ID. Ours to choose; only has to be unique per window.
	idSpotTimer = 1

	// Milliseconds between position updates — roughly 60 per second.
	spotIntervalMs = 16

	// 0 = invisible, 255 = opaque.
	spotAlpha = 160

	// LWA_ALPHA: use the alpha argument. (LWA_COLORKEY, 0x01, is the other mode.)
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setlayeredwindowattributes
	lwaAlpha = 0x00000002
)
