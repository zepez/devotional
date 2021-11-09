package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Devotional struct {
	Id          string    `json:"id"`
	Source      string    `json:"source"`
	Name        string    `json:"name"`
	Html        string    `json:"html"`
	Plain       string    `json:"plain_text"`
	Target_Date string    `json:"target_date"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

func root(w http.ResponseWriter, r *http.Request) {
	devotional := Devotional{}

	devotional.Id = uuid.Must(uuid.NewRandom()).String()
	devotional.Created_at = time.Now()
	devotional.Updated_at = time.Now()
	devotional.Target_Date = getTargetDate(r.URL.Query())
	devotional.Source = GetLink(devotional.Target_Date)

	html, plain, name := Scrape(devotional.Source)

	devotional.Name = name
	devotional.Html = html
	devotional.Plain = plain

	fmt.Fprintf(w, ToJson(devotional))
}

func handleRequests() {
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
