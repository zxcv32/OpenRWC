/*
 * Copyright (c) 2022 OpenRWC.
 *
 * Licensed under GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
 * Everyone is permitted to copy and distribute verbatim copies
 * of this license document, but changing it is not allowed.
 *
 * @author Ashwani Sharma (https://github.com/zxcV32)
 */

package internal

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os/exec"
)

// Set wallpaper on monitor(s)
func XChange(wallpaper string) error {
	_, err := exec.Command("xwallpaper", "--daemon", "--"+viper.GetString("openrwc.util_param"),
		wallpaper).Output()
	if nil != err {
		return err
	}
	log.Infof("X wallpaper applied to all monitors")
	return nil
}
