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
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Load configuration and return the path of config directory
func LoadConfig() string {
	// Load config
	path := "configs"
	file := "config"
	if _, err := os.Stat(path + "/" + file + ".toml"); os.IsNotExist(err) {
		path, file := createDefaultConfig()
		setupViper(path, file, "toml")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err.Error())
		}
		return path
	} else {
		setupViper(path, file, "toml") // Usuall location in project
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err.Error())
		}
		return path
	}
}

func setupViper(dir string, file string, format string) {
	viper.AddConfigPath(dir)
	viper.SetConfigName(file)
	viper.SetConfigType(format)
}

func createDefaultConfig() (string, string) {

	template := `
title = "OpenRWC Configuration"
version = "0.0.1"

[reddit]
subreddits = ["wallpaper", "wallpapers", "Animewallpaper", "AnimeWallpapersSFW", "MinimalWallpaper"]
# one of (hour, day, week, month, year, all)
sort = "hour"
# custom query in addition to openrwc.resolution
query = ""

[openrwc]
# Wallpapert resulutions to search for
# The wallpaper resolution may not always need to match your display resolution
resolutions = [
		"7680x4320", "3840x2160", "1920x1080", "1366x768", "1280x720"
	]

max_attempts = 50
# Number of monitors. Same wallpaper will be set on each monitors
monitors=1

# Nitrogen parameter, One of ("set-auto", "set-centered", "set-scaled", "set-tiled", "set-zoom" , "set-zoom-fill")
nitrogen_param = "set-scaled"

	[openrwc.timeout]
	# s=seconds,m=minutes,h=hours,d=days

	# HTTP call
	call = "15s"
	# Wallpaper refresh frequency. Minimum "2s".
	refresh = "1h"
	# Retry delay if previous query fails. Minimum "2s".
	retry = "5s"
	`
	file := "config"
	home, _ := os.UserHomeDir()
	path := home + "/.config/OpenRWC"
	if _, err := os.Stat(path + "/" + file + ".toml"); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
		f, openError := os.Create(path + "/" + file + ".toml")
		if openError != nil {
			log.Fatal(openError)
		}
		defer f.Close()
		_, writeError := f.WriteString(template)

		if writeError != nil {
			log.Fatal(writeError)
		}
	}
	return path, file
}
