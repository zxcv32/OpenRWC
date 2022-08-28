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
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Set wallpaper on monitor(s)
func FehChange(wallpaper string) error {
	output, err := exec.Command("feh", "--"+viper.GetString("openrwc.feh_bg_setting"),
		wallpaper).Output()
	fmt.Println("Output: " + string(output))
	if nil != err {
		return err
	}
	log.Infof("Feh wallpaper applied to all monitors")
	return nil
}
