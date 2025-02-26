package uuid

import (
	"crypto/rand"
	"time"
)

// NewV7 returns a time-ordered UUID v7 in string format.
//
// The UUID v7 format uses a timestamp with millisecond precision in the most
// significant bits, followed by random data. This provides both uniqueness and
// chronological ordering, making it ideal for database primary keys and
// situations where sorting by creation time is desired.
//
// References:
// - https://uuid7.com/
// - https://www.ietf.org/archive/id/draft-peabody-dispatch-new-uuid-format-04.html#name-uuid-version-7
//
//nolint:lll
func NewV7() (UUID, error) {
	var u UUID

	// Write the timestamp to the first 6 bytes of the UUID.
	timestamp := time.Now().UnixMilli()
	u[0] = byte(timestamp >> 40)
	u[1] = byte(timestamp >> 32)
	u[2] = byte(timestamp >> 24)
	u[3] = byte(timestamp >> 16)
	u[4] = byte(timestamp >> 8)
	u[5] = byte(timestamp)

	// Fill the remaining bytes with random data.
	_, err := rand.Read(u[6:])
	if err != nil {
		// This should never happen.
		return u, err
	}

	// Set the version and variant bits.
	u[6] = (u[6] & 0x0f) | 0x70 // Version: 7 (time-ordered)
	u[8] = (u[8] & 0x3f) | 0x80 // Variant: RFC 4122

	return u, nil
}

// V7Time returns the time of a UUIDv7 string.
//
// If the UUID is not a valid UUIDv7, it returns a zero time and false.
func V7Time(s string) (t time.Time, ok bool) {
	u, err := FromString(s)
	if err != nil {
		return time.Time{}, false
	}

	return u.Time()
}
