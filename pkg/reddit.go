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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Fetch wallpaper URL from reddit
func GetWallpaperUrl(client *http.Client, subreddit string, query string, sort string) ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.reddit.com/r/%s/search.json", subreddit), nil)
	if nil != err {
		return []string{}, err
	}
	req.Header = http.Header{
		"User-Agent": {"OpenRWC - Go"},
	}
	q := req.URL.Query()
	q.Add("q", query)
	q.Add("t", sort)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if nil != err {
		return []string{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return []string{}, errors.New(fmt.Sprintf("Reddit API responded with error respnse status: %d", resp.StatusCode))
	}
	var js interface{}
	apiResponse, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(apiResponse, &js); err != nil {
		return []string{}, err
	}
	defer func() { // Handle panic if API responds with unexpected JSON response
		resp.Body.Close()
		if err := recover(); err != nil {
			log.Errorln("Reddit API did not respond with a wallpaper URL")
		}
	}()
	urls := make([]string, 0)
	children := js.(map[string]interface{})["data"].(map[string]interface{})["children"].([]interface{})
	if nil != children {
		for _, child := range children {
			possibleUrl := child.(map[string]interface{})["data"].(map[string]interface{})["url"]
			validationError := validateWallpaperUrl(possibleUrl.(string))
			if nil != validationError {
				log.Errorln(validationError)
			}
			urls = append(urls, possibleUrl.(string))
		}
	}
	return urls, nil
}

// Try to figure out if the URL possibly points to a wallpaper
func validateWallpaperUrl(possibleUrl string) error {
	errMsg := fmt.Sprintf("Not a wallpaper URL: %s", possibleUrl)
	slashes := strings.Count(possibleUrl, "/")
	if strings.HasSuffix(possibleUrl, "/") {
		return errors.New(errMsg)
	} else if slashes != 3 {
		return errors.New(errMsg)
	}
	return nil
}
