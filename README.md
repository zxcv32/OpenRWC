# OpenRWC

Open Reddit Wallpaper Changer for GNU/Linux

### Requirements
```shell
curl
jq
nitrogen
```

### Steps
1. Place the script that can be located by `$PATH`
   
   `curl <https://raw.githubusercontent.com/zxcV32/OpenRWC/main/openrwc.sh> --output $HOME/bin/openrwc`
2. Make the shell script executable
  
   `chmod +x $HOME/bin/openrwc`
3. Create a cronjob for every 5 minutes with custom log location

   `*/5 * * * * $HOME/bin/openrwc.sh >> $HOME/openrwc.logs`