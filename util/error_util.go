package util

import "log"

func LogError(err error, message string) {
	if err != nil {
		log.Panicf("%s: %s", message, err.Error())
	}
}
