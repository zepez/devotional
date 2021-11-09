package utils

import "encoding/json"

func GenJson(i Devotional) string {
	j, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	return string(j)
}
