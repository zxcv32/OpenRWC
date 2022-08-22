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

	log "github.com/sirupsen/logrus"
)

// Fetch wallpaper URL from reddit
func GetWallpaperUrl(client *http.Client, subreddit string, query string, sort string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.reddit.com/r/%s/search.json", subreddit), nil)
	if nil != err {
		return "", err
	}
	req.Header = http.Header{
		"User-Agent": {"OpenRWC/v0.0.1 - Go"},
	}
	q := req.URL.Query()
	q.Add("q", query)
	q.Add("t", sort)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if nil != err {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("Reddit API responded with error respnse status: %d", resp.StatusCode))
	}
	var js interface{}
	apiResponse, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(apiResponse, &js); err != nil {
		return "", err
	}
	defer func() { // Handle panic if API responds with unexpected JSON response
		resp.Body.Close()
		if err := recover(); err != nil {
			log.Errorln("Reddit API did not responded with a wallpaper URL")
		}
	}()
	url := js.(map[string]interface{})["data"].(map[string]interface{})["children"].([]interface{})[0].(map[string]interface{})["data"].(map[string]interface{})["url"]
	return url.(string), nil
}
