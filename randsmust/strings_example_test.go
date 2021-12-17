package randsmust_test

import (
	"fmt"

	"github.com/jimeh/rands/randsmust"
)

func ExampleBase64() {
	s := randsmust.Base64(16)
	fmt.Println(s) // => rGnZOxJunCd5h+piBpOfDA==
}

func ExampleBase64URL() {
	s := randsmust.Base64URL(16)
	fmt.Println(s) // => NlXKmutou2knLU8q7Hlp5Q
}

func ExampleHex() {
	s := randsmust.Hex(16)
	fmt.Println(s) // => 1013ec67a802be177d3e37f46951e97f
}

func ExampleAlphanumeric() {
	s := randsmust.Alphanumeric(16)
	fmt.Println(s) // => mjT119HdPslVfvUE
}

func ExampleAlphabetic() {
	s := randsmust.Alphabetic(16)
	fmt.Println(s) // => RLaRaTVqcrxvNkiz
}

func ExampleUpper() {
	s := randsmust.Upper(16)
	fmt.Println(s) // => CANJDLMHANPQNXUE
}

func ExampleUpperNumeric() {
	s := randsmust.UpperNumeric(16)
	fmt.Println(s) // => EERZHC96KOIRU9DM
}

func ExampleLower() {
	s := randsmust.Lower(16)
	fmt.Println(s) // => aoybqdwigyezucjy
}

func ExampleLowerNumeric() {
	s := randsmust.LowerNumeric(16)
	fmt.Println(s) // => hs8l2l0750med3g2
}

func ExampleNumeric() {
	s := randsmust.Numeric(16)
	fmt.Println(s) // => 3126402104379869
}

func ExampleString() {
	s := randsmust.String(16, "abcdefABCDEF")
	fmt.Println(s) // => cbdCAbABECaADcaB
}

func ExampleUnicodeString() {
	s := randsmust.UnicodeString(16, []rune("九七二人入八力十下三千上口土夕大"))
	fmt.Println(s) // => 下夕七下千千力入八三力夕千三土七
}

func ExampleDNSLabel() {
	s := randsmust.DNSLabel(16)
	fmt.Println(s) // => urqkt-remuwz5083
}

func ExampleUUID() {
	s := randsmust.UUID()
	fmt.Println(s) // => 5baa35a6-9a46-49b4-91d0-9530173e118d
}
