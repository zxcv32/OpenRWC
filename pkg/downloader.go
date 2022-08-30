/*
 * Copyright (c) 2022 OpenRWC.
 *
 * Licensed under GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
 * Everyone is permitted to copy and distribute verbatim copies
 * of this license document, but changing it is not allowed.
 *
 * @author Ashwani Sharma (https://github.com/zxcV32)
 *
 */

package pkg

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Download image from the given URL to the specified path and return the downloaded file or error/
func Download(urls []string, path string) (string, error) {
	client := &http.Client{} // create a basic downloader client
	for _, url := range urls {
		req, err := http.NewRequest("GET", url, nil)
		if nil != err {
			log.Error(err.Error())
			continue
		}
		req.Header = http.Header{
			"User-Agent": {"OpenRWC - Go"},
		}
		resp, err := client.Do(req)
		if nil != err {
			log.Error(err.Error())
			continue
		}
		if resp.StatusCode != 200 {
			log.Errorf("Received %d HTTP response code from Reddit API", resp.StatusCode)
		}
		defer resp.Body.Close()
		components := strings.Split(url, "i.redd.it/")
		file := path + "/" + components[len(components)-1]
		stat, err := os.Stat(file)
		if errors.Is(err, os.ErrNotExist) {
			if strings.HasSuffix(file, ".png") || strings.HasSuffix(file, ".jpg") || strings.HasSuffix(file, ".jpeg") {
				// This means attempt to download the wallpaper
			} else {
				log.Errorln(err.Error())
				continue
			}
		}
		if nil != stat && stat.Size() > 0 {
			log.Warnf("Wallpaper previousely set: %s", file)
			continue
		}
		newFile, err := os.Create(file)
		if err != nil {
			// big problem
			return "", err
		}
		defer newFile.Close()
		//Write the bytes to the file
		_, err = io.Copy(newFile, resp.Body)
		if err != nil {
			// big problem
			return "", err
		}
		// reached a suitable wallpaper
		return newFile.Name(), nil
	}
	return "", errors.New("No downloadable wallpapers")
}
