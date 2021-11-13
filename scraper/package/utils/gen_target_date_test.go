package utils

import (
	"testing"
	"time"
)

func TestGenTargetDate(t *testing.T) {
	value := GenTargetDate("20060102")
	expect := "20060102"

	if value != expect {
		t.Fatalf(`ERROR: GenTargetDate("20060102") returned %s, expected %s`, value, expect)
	}
}

func TestGenTargetDateEmpty(t *testing.T) {
	value := GenTargetDate("")
	currentTime := time.Now()
	expect := currentTime.Format("20060102")

	if value != expect {
		t.Fatalf(`ERROR: GenTargetDate("") returned %s, expected %s`, value, expect)
	}
}
