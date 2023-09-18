package util

import "log"

func SeperateFatal[T any](val T, err error) T {
	if err != nil {
		log.Fatal(err.Error())
	}

	return val
}
