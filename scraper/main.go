package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	u "scraper/package/utils"
)

func root(w http.ResponseWriter, r *http.Request) {
	devotional := u.Devotional{}

	devotional.Id = uuid.Must(uuid.NewRandom()).String()
	devotional.Created_at = time.Now()
	devotional.Updated_at = time.Now()
	devotional.Target_Date = u.GenTargetDate(r.URL.Query())
	devotional.Source = u.GenLink(devotional.Target_Date)

	html, plain, name := u.Scrape(devotional.Source)

	devotional.Name = name
	devotional.Html = html
	devotional.Plain = plain

	fmt.Fprintf(w, u.GenJson(devotional))
}

func handleRequests() {
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
