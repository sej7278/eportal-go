linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o builds/eportal-go

linux-arm:
	GOOS=linux GOARCH=arm go build -o builds/eportal-go-arm

linux-aarch64:
	GOOS=linux GOARCH=arm64 go build -o builds/eportal-go-aarch64

macos-intel:
	GOOS=darwin GOARCH=amd64 go build -o builds/eportal-go-macos-intel

macos-m2:
	GOOS=darwin GOARCH=arm64 go build -o builds/eportal-go-macos-m2

windows:
	GOOS=windows GOARCH=amd64 go build -o builds/eportal-go.exe

clean:
	rm -f builds/*

all: linux-amd64 linux-arm linux-aarch64 macos-intel macos-m2 windows

.PHONY: all clean linux-amd64 linux-arm linux-aarch64 macos-intel macos-m2 windows
