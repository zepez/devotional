package utils

import (
	"bytes"
	"testing"
)

type StructForTest struct {
	KeyOne string `json:"keyOne"`
	KeyTwo string `json:"keyTwo"`
}

func TestGenJson(t *testing.T) {
	structForTest := StructForTest{}
	structForTest.KeyOne = "hello"
	structForTest.KeyTwo = "world"

	json := GenJson(structForTest)

	expect := []byte{123, 34, 107, 101, 121, 79, 110, 101, 34, 58, 34, 104, 101, 108, 108, 111, 34, 44, 34, 107, 101, 121, 84, 119, 111, 34, 58, 34, 119, 111, 114, 108, 100, 34, 125}

	if bytes.Compare(json, expect) != 0 {
		t.Fatalf(`ERROR: GenJson() incorrect return value`)
	}
}
