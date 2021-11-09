package main

import (
	"net/url"
	"time"
)

func getTargetDate(queries url.Values) string {
	query, queryExists := queries["date"]

	date := ""

	if queryExists {
		date = query[0]
	}

	if !queryExists {
		currentTime := time.Now()
		date = currentTime.Format("20060102")
	}

	return date
}
