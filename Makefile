# Three-part app version; CI passes it from the tag (v1.2.3 -> 1.2.3).
VERSION = 1.0.1
MSIX_DIR := dist/msix
MAKEAPPX ?= makeappx.exe

# Console subsystem: log output visible, Ctrl+C works. For development.
build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

#GUI subsystem: no console window. -s -w strip debug info for a smaller exe.
release:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "-H windowsgui -s -w"

# Lays out the flat folder makeappx packs: exe + manifest + logos side by side.
# The leading space in the sed pattern matches the Identity's Version= only,
# not MinVersion= or MaxVersionTested= on the TargetDeviceFamily line.
msix-stage: release
	rm -rf $(MSIX_DIR)
	mkdir -p $(MSIX_DIR)
	cp spot-on-mouse.exe $(MSIX_DIR)/
	cp -r packaging/Assets $(MSIX_DIR)/
	sed 's/ Version="[0-9.]*"/ Version="$(VERSION).0"/' packaging/AppxManifest.xml > $(MSIX_DIR)/AppxManifest.xml

# Needs the Windows SDK; CI runs makeappx itself on a Windows runner instead.
msix: msix-stage
	$(MAKEAPPX) pack /d $(MSIX_DIR) /p dist/spot-on-mouse-$(VERSION).msix /o
