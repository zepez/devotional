package main

import "encoding/json"

func ToJson(i Devotional) string {
	j, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	return string(j)
}
