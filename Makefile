build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

 # GUI subsystem: no console window. -s -w strip debug info for a smaller exe.
  release:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "-H windowsgui -s -w"
