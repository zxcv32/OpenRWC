# Create executable and system package
SHELL = bash

# Create executable
default:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -tags=nomsgpack -a -o ./build/openrwc-amd64/usr/bin/openrwc ./cmd/
	dpkg-deb --build build/openrwc-amd64
	printf "\n\tsudo dpkg -i build/openrwc-amd64.deb\n"
