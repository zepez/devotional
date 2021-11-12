package utils

import (
	"time"
)

func GenTargetDate(date string) string {

	if len(date) < 1 {
		currentTime := time.Now()
		date = currentTime.Format("20060102")
	}

	return date
}
