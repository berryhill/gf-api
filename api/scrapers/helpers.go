package scrapers

import (
	"strings"
	"time"
)


func TrimSuffix(s, suffix string) string {

	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func TrimPrefix(s string, prefix_length int) string {

	substring := s[prefix_length:]

	return substring
}

func TimeOut(seconds int) error {

	time.Sleep(time.Duration(seconds))

	return nil
}
