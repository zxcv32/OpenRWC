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
	"crypto/tls"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// HTTP client to query wallpapers from Reddit
func TheClientShallIhave() *http.Client {
	var client *http.Client
	if nil == client {
		client = &http.Client{Timeout: getCallTimeout(),
			// Fix for not getting 403 response
			// src: https://www.reddit.com/r/redditdev/comments/uncu00/go_golang_clients_getting_403_blocked_responses/
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{},
			}}
	}
	return client
}

// Get config
func getCallTimeout() time.Duration {
	call, callError := time.ParseDuration(viper.GetString("openrwc.timeout.call"))
	if nil != callError {
		log.Fatalln(callError.Error())
	}
	return call
}
