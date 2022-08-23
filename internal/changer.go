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

package internal

import (
	"errors"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	downloader "github.com/zxcv32/openrwc/pkg"
	reddit "github.com/zxcv32/openrwc/pkg"
)

// Thrown when no wallpaper is set by nitrogen
type NoWallpaperError struct{}

func (m *NoWallpaperError) Error() string {
	return "No wallpaper to change"
}

// Changes wallpaper periodically
func TimedChanger() {
	fails := 0
	for {
		change, err := Change()
		if change {
			// Default wait
			time.Sleep(getRefreshDelay())
		} else {
			if nil != err && errors.Is(err, &NoWallpaperError{}) {
				log.Errorln(err.Error())

			} else {
				fails++
				// Quicly retry to get another wallpaper
				time.Sleep(getRetryDelay())
			}
		}
		if fails > getMaxFails() {
			log.Fatalln("Max fail attempts reached")
		}
	}
}

// Fetch and set wallpaper
func Change() (bool, *NoWallpaperError) {
	path := LoadConfig() // Load config
	for _, subreddit := range getSubreddits() {
		for _, resolution := range getResolutions() {
			path = LoadConfig() // refresh config every time
			searchQuery := strings.Trim(strings.Join([]string{resolution, getQuery()}, " "), " ")
			url, err := reddit.GetWallpaperUrl(TheClientShallIhave(), subreddit, searchQuery, getSort())
			if nil != err {
				log.Errorln(err.Error())
				// Just wait for next iteration if no wallpaper was set
				time.Sleep(getRetryDelay())
				continue
			}
			if len(url) < 1 {
				// No wallpaper URL
				// Just wait for next iteration if no wallpaper was set
				time.Sleep(getRetryDelay())
				continue
			}
			log.Infof("subreddit: %s | %s", subreddit, url)
			log.Infof("config dir: %s", path)
			wallpaper, err := downloader.Download(url, path)
			if nil != err {
				log.Errorln(err.Error())
				// Just wait for next iteration if no wallpaper was set
				time.Sleep(getRetryDelay())
				continue
			}
			log.Infof("wallpaper to set: %s", wallpaper)
			nitrogenError := NitrogenChange(wallpaper)
			if nil != nitrogenError {
				log.Errorf("nitrogen error: %s", nitrogenError.Error())
				// Just wait for next iteration if no wallpaper was set
				time.Sleep(getRetryDelay())
				continue
			}
			return true, nil
		}
	}
	return false, &NoWallpaperError{}
}

// Get config
func getSubreddits() []string {
	subreddits := viper.GetStringSlice("reddit.subreddits")
	if nil == subreddits || len(subreddits) < 1 {
		log.Fatalln("At least on subreddit is required")
	}
	return subreddits
}

// Get config
func getResolutions() []string {
	resolutions := viper.GetStringSlice("openrwc.resolutions")
	if nil == resolutions || len(resolutions) < 1 {
		log.Fatalln("Atleast on screen resolution is required")
	}
	return resolutions
}

// Get config
func getSort() string {
	sort := viper.GetString("reddit.sort")
	if len(strings.TrimSpace(sort)) == 0 {
		log.Fatalln("reddit.sort")
	}
	switch sort {
	// only these are allowed
	case "hour":
	case "day":
	case "week":
	case "month":
	case "year":
	case "all":
	default:
		log.Fatalln("Unknown reddit.sort specified")
	}
	return sort
}

// Get config
func getQuery() string {
	query := viper.GetString("reddit.query")
	return query
}

// Get config
func getRetryDelay() time.Duration {
	retryAfter, retryError := time.ParseDuration(viper.GetString("openrwc.timeout.retry"))
	if nil != retryError {
		log.Fatalln(retryError.Error())
	}
	return retryAfter
}

// Get config
func getRefreshDelay() time.Duration {
	refresh, refreshErr := time.ParseDuration(viper.GetString("openrwc.timeout.refresh"))
	if nil != refreshErr {
		log.Fatalln(refreshErr.Error())
	}
	return refresh
}

// Get config
func getMaxFails() int {
	maxFails := viper.GetInt("max_attempts")
	return maxFails
}