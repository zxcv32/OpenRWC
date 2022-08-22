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
	"log"

	"github.com/spf13/viper"
)

// Load configuration
func LoadConfig() {
	// Load config
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
}
