# OpenRWC

Reddit Wallpaper Changer for GNU/Linux

### Requirements
```shell
curl
jq
nitrogen
```

Configure `nitrogen` to set wallpaper `$HOME/OpenRWC/wp.png`

### Steps
1. Place the script that can be located by `$PATH`
   `curl https://raw.githubusercontent.com/zxcV32/OpenRWC/main/openrwc.sh --output $HOME/bin/openrwc`
2. Make the shell script executable
   `chmod +x $HOME/bin/openrwc`
3. Create a cronjob for every 5 minutes with custom log location
   ```
   crontab -e
   */5 * * * * openrwc >> $HOME/OpenRWC/openrwc.logs
   ```
