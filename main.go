package main

import (
	"fmt"
	"log"
	"net/http"
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

func root(w http.ResponseWriter, r *http.Request) {
	devotional := Devotional{}
	devotional.Id = uuid.Must(uuid.NewRandom()).String()

	currentTime := time.Now()
	devotional.Source = GetLink(currentTime.Format("20060102"))

	html, plain, name := Scrape(devotional.Source)

	devotional.Name = name
	devotional.Html = html
	devotional.Plain = plain

	fmt.Println(ToJson(devotional))
	fmt.Fprintf(w, ToJson(devotional))
}

func handleRequests() {
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
