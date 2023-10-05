linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o builds/eportal-go

linux-arm:
	GOOS=linux GOARCH=arm go build -o builds/eportal-go-arm

linux-aarch64:
	GOOS=linux GOARCH=arm64 go build -o builds/eportal-go-aarch64

macos:
	GOOS=darwin GOARCH=amd64 go build -o builds/eportal-go-macos

windows:
	GOOS=windows GOARCH=amd64 go build -o builds/eportal-go.exe

clean:
	rm -f builds/*

all: clean linux-amd64 linux-arm linux-aarch64 macos windows

.PHONY: all clean linux-amd64 linux-arm linux-aarch64 macos windows
