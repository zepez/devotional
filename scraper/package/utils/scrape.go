package utils

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Scrape(link string) (string, string, string) {
	// request
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	wrapper := doc.Find(".post-body")

	plain := ""
	html := ""
	name := ""

	wrapper.Find("p").Each(func(i int, s *goquery.Selection) {
		if i == 0 || i == 1 {
			paraHtml, err := goquery.OuterHtml(s)
			html = html + paraHtml
			paraPlain := s.Text()
			plain = plain + paraPlain

			if err != nil {
				log.Fatal(err)
			}
		}
	})

	name = wrapper.Find("h2").Text()

	return strings.TrimSpace(html), strings.TrimSpace(plain), strings.TrimSpace(name)
}
