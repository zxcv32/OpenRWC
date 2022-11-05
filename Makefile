# Create executable and system package
SHELL = bash

# Create executable and system pacakge
default:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -tags=nomsgpack -a -o ./build/openrwc-amd64/usr/bin/openrwc ./cmd/

	VERSION=0.0.3 envsubst < build/control > build/openrwc-amd64/DEBIAN/control

	cp build/openrwc@-kde.service build/openrwc-amd64/etc/systemd/system/openrwc@.service
	dpkg-deb --build build/openrwc-amd64
	mv build/openrwc-amd64.deb build/openrwc-kde-amd64.deb

	cp build/openrwc@-nitrogen.service build/openrwc-amd64/etc/systemd/system/openrwc@.service
	echo "Depends: nitrogen (>=1.6.1-2)" >> build/openrwc-amd64/DEBIAN/control
	dpkg-deb --build build/openrwc-amd64
	mv build/openrwc-amd64.deb build/openrwc-nitrogen-amd64.deb

	@echo -e "\n\tsudo dpkg -i build/openrwc-kde-amd64.deb"
	@echo -e "\tsudo dpkg -i build/openrwc-nitrogen-amd64.deb\n"

	@echo -e "Package: openrwc" > build/openrwc-amd64/DEBIAN/control
	@echo -e "[Unit]\nDescription=Reddit Wallpaper Changer for GNU/Linux" > build/openrwc-amd64/etc/systemd/system/openrwc@.service
