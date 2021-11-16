package types

import "time"

// Timestamp wraps arond time.Time and
type Timestamp struct {
	time time.Time
}

// NewTimestamp return a Timestamp object using the current time.
func NewTimestamp() *Timestamp {
	return &Timestamp{time.Now()}
}

// ConvertToTimestamp takes a string, parse it using the RFC3339NanoLenient layout.
// and convert it to a Timestamp object.
func ConvertToTimestamp(timeString string) *Timestamp {

}
