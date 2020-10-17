// Package media provides media writer and filters
package media

import (
	"time"
)

// A Sample contains encoded media and the number of samples in that media
type Sample struct {
	Data      []byte
	Timestamp time.Time
	Duration  time.Duration
}
