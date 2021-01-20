// Package rands provides a suite of functions that use crypto/rand to generate
// cryptographically secure random strings in various formats, as well as ints
// and bytes.
//
// All functions which produce strings from a alphabet of characters uses
// rand.Int() to ensure a uniform distribution of all possible values.
package rands

import "errors"

var errBase = errors.New("rands")
