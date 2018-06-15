package env

import (
	"os"
	"testing"
)

func TestGetEnvWithDefault(t *testing.T) {
	os.Setenv("K1", "V1")

	type testCase struct {
		key      string
		def      string
		expected string
	}

	for _, tc := range []testCase{
		{
			key:      "K1",
			def:      "NOPE",
			expected: "V1",
		}, {
			key:      "K2",
			def:      "NOPE",
			expected: "NOPE",
		}} {
		if actual := GetEnvWithDefault(tc.key, tc.def); actual != tc.expected {
			t.Fatalf("Expected %s but got %s with test %v", tc.expected, actual, tc)
		}
	}
}
