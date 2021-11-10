package utils

import "testing"

func TestGenLink(t *testing.T) {
	date := "20211109"
	link := GenLink(date)
	expect := "https://www.lhm.org/dailydevotions/default.asp?date=20211109"

	if link != expect {
		t.Fatalf(`ERROR: GenLink("20211109") returned %s, expected %s`, link, expect)
	}
}

func TestGenLinkEmpty(t *testing.T) {
	date := ""
	link := GenLink(date)
	expect := "https://www.lhm.org/dailydevotions/default.asp?date="

	if link != expect {
		t.Fatalf(`ERROR: GenLink("20211109") returned %s, expected %s`, link, expect)
	}
}
