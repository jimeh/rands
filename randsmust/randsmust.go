// Package randsmust provides a suite of functions that use crypto/rand to
// generate cryptographically secure random strings in various formats, as well
// as ints and bytes.
//
// All functions which produce strings from a alphabet of characters uses
// rand.Int() to ensure a uniform distribution of all possible values.
//
// randsmust is specifically intended as an alternative to rands for use in
// tests. All functions return a single value, and panic in the event of an
// error. This makes them easy to use when building structs in test cases that
// need random data.
//
// For production code, make sure to use the rands package and check returned
// errors.
package randsmust
