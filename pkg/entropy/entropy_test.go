package entropy

import (
	"regexp"
	"testing"
)

func TestGetEntropy(t *testing.T) {

	cases := []struct {
		desc     string
		mapping  map[*regexp.Regexp]float64
		password string
		expected float64
	}{
		{
			desc: "one rule : only digits",
			mapping: map[*regexp.Regexp]float64{
				regexp.MustCompile("[0-9]"): 3,
			},
			password: "0",
			expected: 3,
		},
		{
			desc: "two rules",
			mapping: map[*regexp.Regexp]float64{
				regexp.MustCompile("[0-9]"): 3,
				regexp.MustCompile("[a-z]"): 6,
			},
			password: "a0",
			expected: 18,
		},
	}

	for _, tc := range cases {
		entropy := GetEntropy(tc.mapping, tc.password)

		if entropy != tc.expected {
			t.Fatalf("%s: expected: %f got: %f for password: %s",
				tc.desc, tc.expected, entropy, tc.password)
		}
	}
}
