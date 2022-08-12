#!/usr/bin/env bash
DIR="$HOME/OpenRWC"
mkdir -p "$DIR"

url=$(curl -s 'https://www.reddit.com/r/wallpaper/search.json?limit=1&q=1920x1080&t=hour' | jq -r '.data|.children|.[0]|.data|.url')
# TODO loop over Subreddits, resolutions and if no wallpaper is found
if [[ "null" == "$url" ]]; then
    echo "no wp found"
else
  echo "URL: $url"
  file="$DIR/wp.$(echo "$url" | awk -F '.' '{print $NF}')"  # path, file, extension
  echo "File: $file"
  curl -s "$url" --output "$file"
  # Update wp
  # Fixme only the wallpaper with extension which was configured in nitrogen is set. Example, if wg.png was set then if wp.jpg is downloaded then it will not be set
  nitrogen --restore
fi
