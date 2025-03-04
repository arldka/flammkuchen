package utils

import (
	"github.com/dustin/go-humanize"
	"time"
)

func RelativeTime(rfc3339Time string) (string, error) {
	// Parse the RFC3339 timestamp
	parsedTime, err := time.Parse(time.RFC3339, rfc3339Time)
	if err != nil {
		return "", err
	}
	// Get the human-readable relative time
	relativeAge := humanize.Time(parsedTime)
	return relativeAge, nil
}
