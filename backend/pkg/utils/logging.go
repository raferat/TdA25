package utils

import "log"

func LogIfErrNotNil(s string, err error) {
	if err != nil {
		log.Printf("%s: %#v\n", s, err)
	}
}
