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
	"os/exec"
	"strings"
)

// XfceChange Set wallpaper on monitor(s)
func XfceChange(wallpaper string) error {
	propsArr, err := exec.Command("xfconf-query", "-c", "xfce4-desktop", "-l").Output()
	if nil != err {
		return err
	}
	props := string(propsArr)
	lines := strings.Split(strings.ReplaceAll(props, "\r\n", "\n"), "\n")
	for _, line := range lines {
		if strings.HasSuffix(line, "/last-image") {
			_, err := exec.Command("xfconf-query", "-c", "xfce4-desktop", "-p", line, "-s",
				wallpaper).Output()
			if nil != err {
				log.Errorf("Error while appling Xfce wallpaper applied to %s", line)
			}
			log.Infof("Xfce wallpaper applied to %s", line)
		}
	}
	return nil
}
