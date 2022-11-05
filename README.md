# OpenRWC

Reddit Wallpaper Changer for GNU/Linux

### Development
```
go run cmd/main.go 
```

## Build
To create a release follow the instructions
1. Set `Version` value in `Makefile`
1. Run `make` to build executable and package using `Makefile` from the project root

### Install System Pacakge
1. Download the latest [release](https://github.com/zxcV32/OpenRWC/releases)

#### KDE
```bash
sudo dpkg -i openrwc-kde-amd64.deb
```

#### Nitrogen
1. Install required dependencies and package
```bash
sudo apt install nitrogen -y
sudo dpkg -i openrwc-nitrogen-amd64.deb
```

### Enable Systemd Service
`sudo systemctl enable openrwc@$USER.service`

### Start Systemd Service
`sudo systemctl start openrwc@$USER.service`

### See Service Logs
`journalctl -fu openrwc@$USER.service`
