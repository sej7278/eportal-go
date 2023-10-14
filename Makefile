$(shell mkdir -p builds)

LDFLAGS=-ldflags "-s -w"

all: linux linux-arm linux-aarch64 linux-ppc64le linux-s390x macos-intel macos windows

linux: # default linux amd64
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o builds/eportal-go

linux-arm: # linux 32-bit armv6 for rpi 0/1/2
	GOOS=linux GOARCH=arm go build $(LDFLAGS) -o builds/eportal-go-arm

linux-aarch64: # linux aarch64 for rpi 3/4/5, graviton, rk3328 etc.
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o builds/eportal-go-aarch64

linux-ppc64le:
	GOOS=linux GOARCH=ppc64le go build $(LDFLAGS) -o builds/eportal-go-ppc64le

linux-s390x:
	GOOS=linux GOARCH=s390x go build $(LDFLAGS) -o builds/eportal-go-s390x

macos: # default macos aarch64
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o builds/eportal-go-macos

macos-intel: # legacy macos amd64
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o builds/eportal-go-macos-intel

windows:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o builds/eportal-go.exe

clean:
	rm -f builds/*

.PHONY: all clean
