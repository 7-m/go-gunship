package utils

import (
	"encoding/json"
	"net/url"
	"strings"
)

// EnsureValue adds the default value to map m if no value exists for key
func EnsureValue(m map[string]map[string]string, key string, defaultValue map[string]string) {
	if _, ok := m[key]; !ok {
		m[key] = defaultValue
	}
}

func MustParseUrl(rawString string) *url.URL {
	urlString, err := url.Parse(rawString)
	if err != nil {
		panic("Failed to parse URL: " + rawString)
	}
	return urlString
}

func Panic(err error, msg string) {
	if err != nil {
		panic(msg + " : [" + err.Error() + "]")
	}
}

func IsTemplate(atom string) bool {
	return strings.HasPrefix(atom, "{") && strings.HasSuffix(atom, "}")
}

func MustMarshal(v interface{}) []byte {
	bytes, err := json.Marshal(v)
	Panic(err, "error while marshalling")
	return bytes
}
