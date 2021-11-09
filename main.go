package main

import (
	"fmt"
	"time"
)

type Devotional struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	// Slug       string    `json:"slug"`
	Name       string    `json:"name"`
	Html       string    `json:"html"`
	Plain      string    `json:"plainText"`
	Created_at time.Time `json:"createdAt"`
	Updated_at time.Time `json:"updatedAt"`
}

func main() {
	devotional := Devotional{}
	devotional.Source = GetLink("20211109")

	html, plain := Scrape(devotional.Source)

	devotional.Html = html
	devotional.Plain = plain

	fmt.Println(devotional)
}
