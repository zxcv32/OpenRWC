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
func NitrogenChange(wallpaper string) error {
	monitors := viper.GetInt("openrwc.monitors")
	for i := 0; i < monitors; i++ {
		_, err := exec.Command("nitrogen", "--"+viper.GetString("openrwc.nitrogen_param"),
			wallpaper, fmt.Sprintf("--head=%d", i)).Output()
		if nil != err {
			return err
		}
		log.Infof("Nitrogen wallpaper applied to monitor: %d", i)
	}
	return nil
}
