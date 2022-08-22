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

package main

import (
	"github.com/sirupsen/logrus"

	openrwc "github.com/zxcv32/openrwc/internal"
)

// Periodically sets wallpaper fetched from reddit
func main() {
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05.999", FullTimestamp: true})
	openrwc.TimedChanger()
}
