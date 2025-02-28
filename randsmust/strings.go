package randsmust

import (
	"github.com/jimeh/rands"
)

// Base64 generates a random base64 encoded string of n number of bytes.
//
// Length of the returned string is about one third greater than the value of n,
// and it may contain characters A-Z, a-z, 0-9, "+", "/", and "=".
func Base64(n int) string {
	r, err := rands.Base64(n)
	if err != nil {
		panic(err)
	}

	return r
}

// Base64URL generates a URL-safe un-padded random base64 encoded string of n
// number of bytes.
//
// Length of the returned string is about one third greater than the value of n,
// and it may contain characters A-Z, a-z, 0-9, "-", and "_".
func Base64URL(n int) string {
	r, err := rands.Base64URL(n)
	if err != nil {
		panic(err)
	}

	return r
}

// Hex generates a random hexadecimal encoded string of n number of bytes.
//
// Length of the returned string is twice the value of n, and it may contain
// characters 0-9 and a-f.
func Hex(n int) string {
	r, err := rands.Hex(n)
	if err != nil {
		panic(err)
	}

	return r
}

// Alphanumeric generates a random alphanumeric string of n length.
//
// The returned string may contain A-Z, a-z, and 0-9.
func Alphanumeric(n int) string {
	r, err := rands.Alphanumeric(n)
	if err != nil {
		panic(err)
	}

	return r
}

// Alphabetic generates a random alphabetic string of n length.
//
// The returned string may contain A-Z, and a-z.
func Alphabetic(n int) string {
	r, err := rands.Alphabetic(n)
	if err != nil {
		panic(err)
	}

	return r
}

// Numeric generates a random numeric string of n length.
//
// The returned string may contain 0-9.
func Numeric(n int) string {
	r, err := rands.Numeric(n)
	if err != nil {
		panic(err)
	}

	return r
}

// Upper generates a random uppercase alphabetic string of n length.
//
// The returned string may contain A-Z.
func Upper(n int) string {
	r, err := rands.Upper(n)
	if err != nil {
		panic(err)
	}

	return r
}

// UpperNumeric generates a random uppercase alphanumeric string of n length.
//
// The returned string may contain A-Z and 0-9.
func UpperNumeric(n int) string {
	r, err := rands.UpperNumeric(n)
	if err != nil {
		panic(err)
	}

	return r
}

// Lower generates a random lowercase alphabetic string of n length.
//
// The returned string may contain a-z.
func Lower(n int) string {
	r, err := rands.Lower(n)
	if err != nil {
		panic(err)
	}

	return r
}

// LowerNumeric generates a random lowercase alphanumeric string of n length.
//
// The returned string may contain A-Z and 0-9.
func LowerNumeric(n int) string {
	r, err := rands.LowerNumeric(n)
	if err != nil {
		panic(err)
	}

	return r
}

// String generates a random string of n length using the given ASCII alphabet.
//
// The specified alphabet determines what characters are used in the returned
// random string. The alphabet can only contain ASCII characters, use
// UnicodeString() if you need a alphabet with Unicode characters.
func String(n int, alphabet string) string {
	r, err := rands.String(n, alphabet)
	if err != nil {
		panic(err)
	}

	return r
}

// UnicodeString generates a random string of n length using the given Unicode
// alphabet.
//
// The specified alphabet determines what characters are used in the returned
// random string. The length of the returned string will be n or greater
// depending on the byte-length of characters which were randomly selected from
// the alphabet.
func UnicodeString(n int, alphabet []rune) string {
	r, err := rands.UnicodeString(n, alphabet)
	if err != nil {
		panic(err)
	}

	return r
}

// DNSLabel returns a random string of n length in a DNS label compliant format
// as defined in RFC 1035, section 2.3.1.
//
// It also adheres to RFC 5891, section 4.2.3.1.
//
// In summary, the generated random string will:
//
//   - be between 1 and 63 characters in length, other n values returns a error
//   - first character will be one of a-z
//   - last character will be one of a-z or 0-9
//   - in-between first and last characters consist of a-z, 0-9, or "-"
//   - potentially contain two or more consecutive "-", except the 3rd and 4th
//     characters, as that would violate RFC 5891, section 4.2.3.1.
func DNSLabel(n int) string {
	r, err := rands.DNSLabel(n)
	if err != nil {
		panic(err)
	}

	return r
}

// UUIDv4 returns a random UUID v4 in string format as defined by RFC 4122,
// section 4.4.
func UUID() string {
	r, err := rands.UUID()
	if err != nil {
		panic(err)
	}

	return r
}

// UUIDv7 returns a time-ordered UUID v7 in string format.
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
func UUIDv7() string {
	r, err := rands.UUIDv7()
	if err != nil {
		panic(err)
	}

	return r
}
