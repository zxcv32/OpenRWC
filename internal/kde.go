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
	"os"
	"os/exec"
)

// KdeChange Set wallpaper on monitor(s)
// path to the script
func KdeChange(path string, wallpaper string) error {
	// https://www.reddit.com/r/kde/comments/65pmhj/change_wallpaper_from_terminal/
	script := `#!/usr/bin/env sh
qdbus org.kde.plasmashell /PlasmaShell org.kde.PlasmaShell.evaluateScript "var allDesktops = desktops();
                                                                           for (i=0;i<allDesktops.length;i++) {{
                                                                               d = allDesktops[i];
                                                                               d.wallpaperPlugin = \"org.kde.image\";
                                                                               d.currentConfigGroup = Array(\"Wallpaper\",
                                                                                                            \"org.kde.image\",
                                                                                                            \"General\");
                                                                               d.writeConfig(\"Image\", \"$1\")
                                                                           }}
"
`
	file := "kde.sh"
	if _, err := os.Stat(path + "/" + file); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
		f, openError := os.Create(path + "/" + file)
		if openError != nil {
			log.Fatal(openError)
		}
		defer f.Close()
		_, writeError := f.WriteString(script)

		if writeError != nil {
			log.Fatal(writeError)
		}
		err = os.Chmod(path+"/"+file, 0775)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("Shell script %s created at %s", file, path)
	}
	_, err := exec.Command(path+"/"+file, wallpaper).Output()
	if nil != err {
		return err
	}
	log.Infof("KDE wallpaper applied to all monitors")
	return nil
}
