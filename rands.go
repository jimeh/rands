// Package rands provides a suite of functions that use crypto/rand to generate
// cryptographically secure random strings in various formats, as well as ints
// and bytes.
//
// All functions which produce strings from a alphabet of characters uses
// rand.Int() to ensure a uniform distribution of all possible values.
//
// rands is intended for use in production code where random data generation is
// required. All functions have a error return value, which should be
// checked.
//
// For tests there is the randsmust package, which has all the same functions
// but with single return values, and they panic in the event of an error.
package rands

import "errors"

var Err = errors.New("rands")
