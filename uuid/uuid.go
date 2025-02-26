// Package uuid provides a UUID type and associated utilities.
package uuid

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

var (
	Err              = errors.New("uuid")
	ErrInvalidLength = fmt.Errorf("%w: invalid length", Err)
)

const (
	hyphen = '-'
)

// UUID represents a Universally Unique Identifier (UUID).
// It is implemented as a 16-byte array.
type UUID [16]byte

// String returns the string representation of the UUID,
// formatted according to RFC 4122 (8-4-4-4-12 hex digits separated by hyphens).
func (u UUID) String() string {
	dst := make([]byte, 36)

	hex.Encode(dst[0:8], u[0:4])
	dst[8] = hyphen
	hex.Encode(dst[9:13], u[4:6])
	dst[13] = hyphen
	hex.Encode(dst[14:18], u[6:8])
	dst[18] = hyphen
	hex.Encode(dst[19:23], u[8:10])
	dst[23] = hyphen
	hex.Encode(dst[24:], u[10:])

	return string(dst)
}

// FromBytes creates a UUID from a byte slice.
//
// If the slice isn't exactly 16 bytes, it returns an empty UUID.
func FromBytes(b []byte) (UUID, error) {
	var u UUID
	if len(b) != 16 {
		return u, ErrInvalidLength
	}

	copy(u[:], b)

	return u, nil
}

// FromString creates a UUID from a string.
//
// If the string isn't exactly 36 characters, it returns an empty UUID.
func FromString(s string) (UUID, error) {
	if len(s) != 36 {
		return UUID{}, ErrInvalidLength
	}

	raw := strings.ReplaceAll(s, "-", "")

	u := UUID{}
	_, err := hex.Decode(u[:], []byte(raw))
	if err != nil {
		return UUID{}, err
	}

	return u, nil
}

// Time returns the timestamp of the UUID if it's a version 7 (time-ordered)
// UUID. Otherwise, it returns the zero time.
func (u UUID) Time() (t time.Time, ok bool) {
	if u.Version() != 7 {
		return time.Time{}, false
	}

	// Extract the timestamp from the UUID.
	// For UUIDv7, only the first 6 bytes contain the timestamp in milliseconds
	timestamp := uint64(u[0])<<40 | uint64(u[1])<<32 | uint64(u[2])<<24 |
		uint64(u[3])<<16 | uint64(u[4])<<8 | uint64(u[5])

	if timestamp > math.MaxInt64 {
		// This shouldn't happen until year 292,272,993.
		return time.Time{}, false
	}

	return time.UnixMilli(int64(timestamp)), true
}

// Version returns the version of the UUID.
func (u UUID) Version() int {
	// The version is stored in the 4 most significant bits of byte 6
	return int(u[6] >> 4)
}
