package utils

import (
	"os"
	"testing"
)

func TestGenEnv(t *testing.T) {
	os.Setenv("PORT", "8080")
	value := GetEnvUtil("PORT", "8080")
	expect := "8080"

	if value != expect {
		t.Fatalf(`ERROR: GetEnvUtil("PORT", "8080") returned %s, expected %s`, value, expect)
	}
}

func TestGenEnvFallbackUnused(t *testing.T) {
	os.Setenv("PORT", "8080")
	value := GetEnvUtil("PORT", "3000")
	expect := "8080"

	if value != expect {
		t.Fatalf(`ERROR: GetEnvUtil("PORT", "3000") returned %s, expected %s`, value, expect)
	}
}

func TestGenEnvFallbackUsed(t *testing.T) {
	value := GetEnvUtil("PORT", "8080")
	expect := "8080"

	if value != expect {
		t.Fatalf(`ERROR: GetEnvUtil("PORT", "8081") returned %s, expected %s`, value, expect)
	}
}
