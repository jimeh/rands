package uuid

import (
	"crypto/rand"
)

// NewRandom returns a random UUID v4 in string format as defined by RFC 4122,
// section 4.4.
func NewRandom() (UUID, error) {
	var u UUID

	// Fill the entire UUID with random bytes.
	_, err := rand.Read(u[:])
	if err != nil {
		// This should never happen.
		return u, err
	}

	// Set the version and variant bits.
	u[6] = (u[6] & 0x0f) | 0x40 // Version: 4 (random)
	u[8] = (u[8] & 0x3f) | 0x80 // Variant: RFC 4122

	return u, nil
}
