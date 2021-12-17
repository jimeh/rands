package rands

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"unicode"
)

const (
	upperChars        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars        = "abcdefghijklmnopqrstuvwxyz"
	numericChars      = "0123456789"
	lowerNumericChars = lowerChars + numericChars
	upperNumericChars = upperChars + numericChars
	alphabeticChars   = upperChars + lowerChars
	alphanumericChars = alphabeticChars + numericChars
	dnsLabelChars     = lowerNumericChars + "-"
	uuidHyphen        = byte('-')
)

var (
	ErrNonASCIIAlphabet = fmt.Errorf(
		"%w: alphabet contains non-ASCII characters", Err,
	)

	ErrDNSLabelLength = fmt.Errorf(
		"%w: DNS labels must be between 1 and 63 characters in length", Err,
	)
)

// Base64 generates a random base64 encoded string of n number of bytes.
//
// Length of the returned string is about one third greater than the value of n,
// and it may contain characters A-Z, a-z, 0-9, "+", "/", and "=".
func Base64(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

// Base64URL generates a URL-safe un-padded random base64 encoded string of n
// number of bytes.
//
// Length of the returned string is about one third greater than the value of n,
// and it may contain characters A-Z, a-z, 0-9, "-", and "_".
func Base64URL(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

// Hex generates a random hexadecimal encoded string of n number of bytes.
//
// Length of the returned string is twice the value of n, and it may contain
// characters 0-9 and a-f.
func Hex(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

// Alphanumeric generates a random alphanumeric string of n length.
//
// The returned string may contain A-Z, a-z, and 0-9.
func Alphanumeric(n int) (string, error) {
	return String(n, alphanumericChars)
}

// Alphabetic generates a random alphabetic string of n length.
//
// The returned string may contain A-Z, and a-z.
func Alphabetic(n int) (string, error) {
	return String(n, alphabeticChars)
}

// Numeric generates a random numeric string of n length.
//
// The returned string may contain 0-9.
func Numeric(n int) (string, error) {
	return String(n, numericChars)
}

// Upper generates a random uppercase alphabetic string of n length.
//
// The returned string may contain A-Z.
func Upper(n int) (string, error) {
	return String(n, upperChars)
}

// UpperNumeric generates a random uppercase alphanumeric string of n length.
//
// The returned string may contain A-Z and 0-9.
func UpperNumeric(n int) (string, error) {
	return String(n, upperNumericChars)
}

// Lower generates a random lowercase alphabetic string of n length.
//
// The returned string may contain a-z.
func Lower(n int) (string, error) {
	return String(n, lowerChars)
}

// LowerNumeric generates a random lowercase alphanumeric string of n length.
//
// The returned string may contain A-Z and 0-9.
func LowerNumeric(n int) (string, error) {
	return String(n, lowerNumericChars)
}

// String generates a random string of n length using the given ASCII alphabet.
//
// The specified alphabet determines what characters are used in the returned
// random string. The alphabet can only contain ASCII characters, use
// UnicodeString() if you need a alphabet with Unicode characters.
func String(n int, alphabet string) (string, error) {
	if !isASCII(alphabet) {
		return "", ErrNonASCIIAlphabet
	}

	l := big.NewInt(int64(len(alphabet)))
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		index, err := rand.Int(rand.Reader, l)
		if err != nil {
			return "", err
		}
		b[i] = alphabet[index.Int64()]
	}

	return string(b), nil
}

// UnicodeString generates a random string of n length using the given Unicode
// alphabet.
//
// The specified alphabet determines what characters are used in the returned
// random string. The length of the returned string will be n or greater
// depending on the byte-length of characters which were randomly selected from
// the alphabet.
func UnicodeString(n int, alphabet []rune) (string, error) {
	l := big.NewInt(int64(len(alphabet)))
	b := make([]rune, n)
	for i := 0; i < n; i++ {
		index, err := rand.Int(rand.Reader, l)
		if err != nil {
			return "", err
		}
		b[i] = alphabet[index.Int64()]
	}

	return string(b), nil
}

// DNSLabel returns a random string of n length in a DNS label compliant format
// as defined in RFC 1035, section 2.3.1.
//
// It also adheres to RFC 5891, section 4.2.3.1.
//
// In summary, the generated random string will:
//
//  - be between 1 and 63 characters in length, other n values returns a error
//  - first character will be one of a-z
//  - last character will be one of a-z or 0-9
//  - in-between first and last characters consist of a-z, 0-9, or "-"
//  - potentially contain two or more consecutive "-", except the 3rd and 4th
//    characters, as that would violate RFC 5891, section 4.2.3.1.
func DNSLabel(n int) (string, error) {
	switch {
	case n < 1 || n > 63:
		return "", ErrDNSLabelLength
	case n == 1:
		return String(1, lowerChars)
	default:
		// First character of a DNS label allows only a-z characters.
		head, err := String(1, lowerChars)
		if err != nil {
			return "", err
		}

		// Last character of a DNS label allows only a-z and 0-9 characters.
		tail, err := String(1, lowerNumericChars)
		if err != nil {
			return "", err
		}

		if n < 3 {
			return head + tail, nil
		}

		// The middle of a DNS label allows only a-z, 0-9, and "-" characters.
		bodyLen := n - 2
		body := make([]byte, bodyLen)
		var last byte
		var l *big.Int

		for i := 0; i < bodyLen; i++ {
			// Prevent two consecutive hyphens characters in positions 3 and 4,
			// in accordance RFC 5891, section 4.2.3.1, Hyphen Restrictions:
			// https://tools.ietf.org/html/rfc5891#section-4.2.3.1
			if i == 2 && last == byte(45) {
				l = big.NewInt(int64(len(lowerNumericChars)))
			} else {
				l = big.NewInt(int64(len(dnsLabelChars)))
			}

			index, err := rand.Int(rand.Reader, l)
			if err != nil {
				return "", err
			}

			if i == 2 && last == byte(45) {
				last = lowerNumericChars[index.Int64()]
			} else {
				last = dnsLabelChars[index.Int64()]
			}

			body[i] = last
		}

		return head + string(body) + tail, nil
	}
}

// UUID returns a random UUID v4 in string format as defined by RFC 4122,
// section 4.4.
func UUID() (string, error) {
	b, err := Bytes(16)
	if err != nil {
		return "", err
	}

	b[6] = (b[6] & 0x0f) | 0x40 // Version: 4 (random)
	b[8] = (b[8] & 0x3f) | 0x80 // Variant: RFC 4122

	// Construct a UUID v4 string according to RFC 4122 specifications.
	dst := make([]byte, 36)
	hex.Encode(dst[0:8], b[0:4]) // time-low
	dst[8] = uuidHyphen
	hex.Encode(dst[9:13], b[4:6]) // time-mid
	dst[13] = uuidHyphen
	hex.Encode(dst[14:18], b[6:8]) // time-high-and-version
	dst[18] = uuidHyphen
	hex.Encode(dst[19:23], b[8:10]) // clock-seq-and-reserved, clock-seq-low
	dst[23] = uuidHyphen
	hex.Encode(dst[24:], b[10:]) // node

	return string(dst), nil
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}

	return true
}
