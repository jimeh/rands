package rands_test

import (
	"fmt"
	"log"

	"github.com/jimeh/rands"
)

func ExampleBase64() {
	s, err := rands.Base64(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => nYQLhIYTqh8oH/W4hZuXMQ==
}

func ExampleBase64URL() {
	s, err := rands.Base64URL(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => zI_zrc1l0uPT4MxncR6e5w
}

func ExampleHex() {
	s, err := rands.Hex(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => b59e8977a13f3c030bd2ea1002ec8081
}

func ExampleAlphanumeric() {
	s, err := rands.Alphanumeric(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => EgPieCBO7MuWhHtj
}

func ExampleAlphabetic() {
	s, err := rands.Alphabetic(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => VzcovEqvMRBWUtQC
}

func ExampleUpper() {
	s, err := rands.Upper(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => MCZEGPWGYKNUEDCK
}

func ExampleUpperNumeric() {
	s, err := rands.UpperNumeric(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => 6LLPBBUW77B26X2X
}

func ExampleLower() {
	s, err := rands.Lower(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => dhoqhrqljadsztaa
}

func ExampleLowerNumeric() {
	s, err := rands.LowerNumeric(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => th1z1b1d24l5h8pu
}

func ExampleNumeric() {
	s, err := rands.Numeric(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => 3378802228987741
}

func ExampleString() {
	s, err := rands.String(16, "abcdefABCDEF")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => BAFffADaadeeacfa
}

func ExampleUnicodeString() {
	s, err := rands.UnicodeString(16, []rune("九七二人入八力十下三千上口土夕大"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => 八三口上土土七入力夕人力下三上力
}

func ExampleDNSLabel() {
	s, err := rands.DNSLabel(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => ab-sbh5q0gfb6sqo
}

func ExampleUUID() {
	s, err := rands.UUID()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // => 6a1c4f65-d5d6-4a28-aa51-eaa94fa7ad4a
}
