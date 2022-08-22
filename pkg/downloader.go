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

// Download image from the given URL and return the downloaded file or error/
func Download(url string) (string, error) {
	client := &http.Client{} // create a basic downloader client
	resp, err := client.Get(url)
	if nil != err {
		log.Error(err)
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New("received non 200 response code")
	}
	defer resp.Body.Close()
	components := strings.Split(url, "i.redd.it/")
	filename := components[len(components)-1]
	stat, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		if strings.HasSuffix(filename, ".png") || strings.HasSuffix(filename, ".jpg") || strings.HasSuffix(filename, ".jpeg") {
			// This means attempt to download the wallpaper
		} else {
			return "", err
		}
	}
	if nil != stat && stat.Size() > 0 {
		// do not return a wallpaper previousely set
		return "", errors.New("Wallpaper exists: " + filename)
	}
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()
	//Write the bytes to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}
	return file.Name(), nil
}
