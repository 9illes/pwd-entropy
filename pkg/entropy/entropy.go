package entropy

import (
	"regexp"
)

// DefaultRules is a mapping from char namespace to entropy
var DefaultRules = map[string]float64{
	"[0-9]":                                  3.322,
	"[a-z]":                                  4.7,
	"[A-Z]":                                  4.7,
	"[\x20-\x2F\x3A-\x40\x5B-\x60\x7B-\x7E]": 5.3, // ASCII printables specials chars
}

// GetEntropy returns entropy of a password (~bits)
func GetEntropy(mapping map[*regexp.Regexp]float64, password string) float64 {

	entropy := 0.0

	for re, f := range mapping {
		if re.Match([]byte(password)) {
			entropy += f
		}
	}

	return entropy * float64(len(password))
}

// NewNamespaceToEntropy returns a mapping from namespace to entropy
func NewNamespaceToEntropy(mapping map[string]float64) map[*regexp.Regexp]float64 {
	namespace := map[*regexp.Regexp]float64{}

	for reStr, f := range mapping {
		r := regexp.MustCompile(reStr)
		namespace[r] = f
	}

	return namespace
}
