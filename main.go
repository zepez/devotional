package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Devotional struct {
	Id         string    `json:"id"`
	Source     string    `json:"source"`
	Name       string    `json:"name"`
	Html       string    `json:"html"`
	Plain      string    `json:"plainText"`
	Created_at time.Time `json:"createdAt"`
	Updated_at time.Time `json:"updatedAt"`
}

func main() {
	devotional := Devotional{}
	devotional.Id = uuid.Must(uuid.NewRandom()).String()
	devotional.Source = GetLink("20211109")

	html, plain, name := Scrape(devotional.Source)

	devotional.Name = name
	devotional.Html = html
	devotional.Plain = plain

	fmt.Println(ToJson(devotional))
}
