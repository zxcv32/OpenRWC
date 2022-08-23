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
func Download(url string, path string) (string, error) {
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
	file := path + "/" + components[len(components)-1]
	log.Infof("Download path: %s", file)
	stat, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		if strings.HasSuffix(file, ".png") || strings.HasSuffix(file, ".jpg") || strings.HasSuffix(file, ".jpeg") {
			// This means attempt to download the wallpaper
		} else {
			return "", err
		}
	}
	if nil != stat && stat.Size() > 0 {
		// do not return a wallpaper previousely set
		return "", errors.New("Wallpaper exists: " + file)
	}
	newFile, err := os.Create(file)
	if err != nil {
		return "", err
	}

	defer newFile.Close()
	//Write the bytes to the file
	_, err = io.Copy(newFile, resp.Body)
	if err != nil {
		return "", err
	}
	return newFile.Name(), nil
}
