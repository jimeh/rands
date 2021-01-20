package rands_test

import (
	"fmt"

	"github.com/jimeh/rands"
)

func ExampleBase64() {
	s, _ := rands.Base64(16)
	fmt.Println(s) // => CYxqEdUB1Rzno3SyZu2g/g==
}

func ExampleBase64URL() {
	s, _ := rands.Base64URL(16)
	fmt.Println(s) // => zlqw9aFqcFggbk2asn3_aQ
}

func ExampleHex() {
	s, _ := rands.Hex(16)
	fmt.Println(s) // => 956e2ec9e7f19ddd58bb935826926531
}

func ExampleAlphanumeric() {
	s, _ := rands.Alphanumeric(16)
	fmt.Println(s) // => Fvk1PkrmG5crgOjT
}

func ExampleAlphabetic() {
	s, _ := rands.Alphabetic(16)
	fmt.Println(s) // => XEJIzcZufHkuUmRM
}

func ExampleUpper() {
	s, _ := rands.Upper(16)
	fmt.Println(s) // => UMAGAFPPNDRGLUPZ
}

func ExampleUpperNumeric() {
	s, _ := rands.UpperNumeric(16)
	fmt.Println(s) // => DF0CQS0TK9CPUO3E
}

func ExampleLower() {
	s, _ := rands.Lower(16)
	fmt.Println(s) // => ocsmggykzrxzfwgt
}

func ExampleLowerNumeric() {
	s, _ := rands.LowerNumeric(16)
	fmt.Println(s) // => rwlv7a1p7klqffs5
}

func ExampleNumeric() {
	s, _ := rands.Numeric(16)
	fmt.Println(s) // => 9403373143598295
}

func ExampleString() {
	s, _ := rands.String(16, "abcdefABCDEF")
	fmt.Println(s) // => adCDCaDEdeffeDeb
}

func ExampleUnicodeString() {
	s, _ := rands.UnicodeString(16, []rune("九七二人入八力十下三千上口土夕大"))
	fmt.Println(s) // => 下下口九力下土夕下土八上二夕大三
}

func ExampleDNSLabel() {
	s, _ := rands.DNSLabel(16)
	fmt.Println(s) // => z0ij9o8qkbs0ru-h
}
