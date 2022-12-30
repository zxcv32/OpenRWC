# OpenRWC

Reddit Wallpaper Changer for GNU/Linux

## Install System Package
1. Download the latest [release](https://github.com/zxcV32/OpenRWC/releases)
2. Install the relevant package depending upon your DE or WM
   ```shell
   sudo --preserve-env=HOME,USER dpkg -i build/openrwc-kde-amd64.deb
   sudo --preserve-env=HOME,USER dpkg -i build/openrwc-nitrogen-amd64.deb
   sudo --preserve-env=HOME,USER dpkg -i build/openrwc-x-amd64.deb
   sudo --preserve-env=HOME,USER dpkg -i build/openrwc-xfce-amd64.deb
   ```

## Enable & Start Systemd Service
```shell
systemctl --user enable openrwc.service
systemctl --user start openrwc.service
```

## See Service Logs
```shell
journalctl -f --user-unit openrwc.service
```

## Development
```
go run cmd/main.go 
```

## Build
To create a release follow the instructions
1. Set `Version` value in `Makefile`
2. Run `make` to build executable and package using `Makefile` from the project root

### Additional Info
1. Configuration file will be created after the first run at `~/.config/OpenRWC/config.toml`
2. Wallpapers are downloaded at `~/.config/OpenRWC/`
3. Systemd Service is created in the user's home directory: `~/.config/systemd/user/openrwc.service`