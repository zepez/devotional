package utils

import "encoding/json"

func GenJson(i interface{}) []byte {
	j, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	return j
}
