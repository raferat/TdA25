package utils

import (
	"fmt"
	"log"
	"time"
)

func LogIfErrNotNil(s string, err error) {
	if err != nil {
		log.Printf("%s: %#v\n", s, err)
	}
}

func datetimeFormat(t time.Time) string {
	year, month, day := t.Date()
	hour, minute, second := t.Clock()

	return fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
}
