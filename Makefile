# Create executable and system pacakge
	
# Create executable
default:
	CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -tags=nomsgpack -a -o ./build/openrwc-amd64/usr/bin/openrwc ./cmd/
	dpkg-deb --build build/openrwc-amd64
