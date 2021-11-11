package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	def "backend/package/definitions"
)

func GetScrapedUtil() def.Devotional {

	res, err := http.Get("http://localhost:8081?date=20211105")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result def.Devotional
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("error:", err)
	}

	return result
}
