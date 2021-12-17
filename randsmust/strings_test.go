package randsmust

import (
	"encoding/base64"
	"encoding/hex"
	"regexp"
	"strings"
	"testing"

	"github.com/jimeh/rands"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHex(t *testing.T) {
	t.Parallel()

	allowed := "0123456789abcdef"

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Hex(tt.n)

			assert.Len(t, got, tt.n*2)
			assertAllowedChars(t, allowed, got)
		})
	}
}

func TestBase64(t *testing.T) {
	t.Parallel()

	allowed := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" +
		"0123456789+/="

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Base64(tt.n)

			b, err := base64.StdEncoding.DecodeString(got)
			require.NoError(t, err)

			assert.Len(t, b, tt.n)
			assertAllowedChars(t, allowed, got)
		})
	}
}

func TestBase64URL(t *testing.T) {
	t.Parallel()

	allowed := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" +
		"0123456789-_"

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Base64URL(tt.n)

			b, err := base64.RawURLEncoding.DecodeString(got)
			require.NoError(t, err)

			assert.Len(t, b, tt.n)
			assertAllowedChars(t, allowed, got)
		})
	}
}

func TestAlphanumeric(t *testing.T) {
	t.Parallel()

	allowed := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Alphanumeric(tt.n)

			assert.Len(t, got, tt.n)
			assertAllowedChars(t, allowed, got)
		})
	}
}

func TestAlphabetic(t *testing.T) {
	t.Parallel()

	allowed := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Alphabetic(tt.n)

			assert.Len(t, got, tt.n)
			assertAllowedChars(t, allowed, got)
		})
	}
}

func TestNumeric(t *testing.T) {
	t.Parallel()

	allowed := "0123456789"

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Numeric(tt.n)

			assert.Len(t, got, tt.n)
			assertAllowedChars(t, allowed, got)
		})
	}
}

func TestUpper(t *testing.T) {
	t.Parallel()

	allowed := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Upper(tt.n)

			assert.Len(t, got, tt.n)
			assertAllowedChars(t, allowed, got)
		})
	}
}

func TestUpperNumeric(t *testing.T) {
	t.Parallel()

	allowed := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := UpperNumeric(tt.n)

			assert.Len(t, got, tt.n)
			assertAllowedChars(t, allowed, got)
		})
	}
}

func TestLower(t *testing.T) {
	t.Parallel()

	allowed := "abcdefghijklmnopqrstuvwxyz"

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Lower(tt.n)

			assert.Len(t, got, tt.n)
			assertAllowedChars(t, allowed, got)
		})
	}
}

func TestLowerNumeric(t *testing.T) {
	t.Parallel()

	allowed := "abcdefghijklmnopqrstuvwxyz0123456789"

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := LowerNumeric(tt.n)

			assert.Len(t, got, tt.n)
			assertAllowedChars(t, allowed, got)
		})
	}
}

var stringTestCases = []struct {
	name       string
	n          int
	alphabet   string
	panicErrIs error
	panicStr   string
}{
	{
		name: "greek",
		n:    32,
		alphabet: "αβγδεζηθικλμνξοπρστυφχψωςΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΩ" +
			"άίόύώέϊϋΐΰΆΈΌΏΎΊ",
		panicErrIs: rands.ErrNonASCIIAlphabet,
		panicStr:   "rands: alphabet contains non-ASCII characters",
	},
	{
		name: "chinese",
		n:    32,
		alphabet: "的一是不了人我在有他这为之大来以个中上们到说国和地也子" +
			"时道出而要于就下得可你年生",
		panicErrIs: rands.ErrNonASCIIAlphabet,
		panicStr:   "rands: alphabet contains non-ASCII characters",
	},
	{
		name: "japanese",
		n:    32,
		alphabet: "一九七二人入八力十下三千上口土夕大女子小山川五天中六円" +
			"手文日月木水火犬王正出本右四",
		panicErrIs: rands.ErrNonASCIIAlphabet,
		panicStr:   "rands: alphabet contains non-ASCII characters",
	},
	{
		name:     "n=0",
		n:        0,
		alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-",
	},
	{
		name:     "n=1",
		n:        1,
		alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-",
	},
	{
		name:     "n=2",
		n:        2,
		alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-",
	},
	{
		name:     "n=7",
		n:        7,
		alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-",
	},
	{
		name:     "n=8",
		n:        8,
		alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-",
	},
	{
		name:     "n=16",
		n:        16,
		alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-",
	},
	{
		name:     "n=32",
		n:        32,
		alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-",
	},
	{
		name:     "n=128",
		n:        128,
		alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-",
	},
	{
		name:     "n=1024",
		n:        1024,
		alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-",
	},
	{
		name:     "n=409600",
		n:        409600,
		alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-",
	},
	{
		name:     "n=2394345",
		n:        2394345,
		alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-",
	},
	{
		name:     "uppercase",
		n:        16,
		alphabet: "ABCDEFGHJKMNPRSTUVWXYZ",
	},
	{
		name:     "lowercase",
		n:        16,
		alphabet: "abcdefghjkmnprstuvwxyz",
	},
}

func TestString(t *testing.T) {
	t.Parallel()

	for _, tt := range stringTestCases {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			p := recoverPanic(func() {
				got = String(tt.n, tt.alphabet)
			})

			if tt.panicErrIs == nil || tt.panicStr == "" {
				assert.Len(t, []rune(got), tt.n)
				assertAllowedChars(t, tt.alphabet, got)
			}

			if tt.panicErrIs != nil {
				assert.ErrorIs(t, p.(error), tt.panicErrIs)
			}

			if tt.panicStr != "" {
				assert.EqualError(t, p.(error), tt.panicStr)
			}
		})
	}
}

var unicodeStringTestCases = []struct {
	name     string
	n        int
	alphabet string
}{
	{
		name: "n=0",
		n:    0,
		alphabet: "αβγδεζηθικλμν时道出而要于就下得可你年" +
			"手文日月木水火犬王正出本右",
	},
	{
		name: "n=1",
		n:    1,
		alphabet: "αβγδεζηθικλμν时道出而要于就下得可你年" +
			"手文日月木水火犬王正出本右",
	},
	{
		name: "n=2",
		n:    2,
		alphabet: "αβγδεζηθικλμν时道出而要于就下得可你年" +
			"手文日月木水火犬王正出本右",
	},
	{
		name: "n=7",
		n:    7,
		alphabet: "αβγδεζηθικλμν时道出而要于就下得可你年" +
			"手文日月木水火犬王正出本右",
	},
	{
		name: "n=8",
		n:    8,
		alphabet: "αβγδεζηθικλμν时道出而要于就下得可你年" +
			"手文日月木水火犬王正出本右",
	},
	{
		name: "n=16",
		n:    16,
		alphabet: "αβγδεζηθικλμν时道出而要于就下得可你年" +
			"手文日月木水火犬王正出本右",
	},
	{
		name: "n=32",
		n:    32,
		alphabet: "αβγδεζηθικλμν时道出而要于就下得可你年" +
			"手文日月木水火犬王正出本右",
	},
	{
		name: "n=128",
		n:    128,
		alphabet: "αβγδεζηθικλμν时道出而要于就下得可你年" +
			"手文日月木水火犬王正出本右",
	},
	{
		name: "n=1024",
		n:    1024,
		alphabet: "αβγδεζηθικλμν时道出而要于就下得可你年" +
			"手文日月木水火犬王正出本右",
	},
	{
		name: "n=409600",
		n:    409600,
		alphabet: "αβγδεζηθικλμν时道出而要于就下得可你年" +
			"手文日月木水火犬王正出本右",
	},
	{
		name: "n=2394345",
		n:    2394345,
		alphabet: "αβγδεζηθικλμν时道出而要于就下得可你年" +
			"手文日月木水火犬王正出本右",
	},
	{
		name:     "latin",
		n:        32,
		alphabet: "ABCDEFGHJKMNPRSTUVWXYZabcdefghjkmnprstuvwxyz",
	},
	{
		name: "greek",
		n:    32,
		alphabet: "αβγδεζηθικλμνξοπρστυφχψωςΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΩ" +
			"άίόύώέϊϋΐΰΆΈΌΏΎΊ",
	},
	{
		name: "chinese",
		n:    32,
		alphabet: "的一是不了人我在有他这为之大来以个中上们到说国和地也子" +
			"时道出而要于就下得可你年生",
	},
	{
		name: "japanese",
		n:    32,
		alphabet: "一九七二人入八力十下三千上口土夕大女子小山川五天中六円" +
			"手文日月木水火犬王正出本右四",
	},
}

func TestUnicodeString(t *testing.T) {
	t.Parallel()

	for _, tt := range unicodeStringTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := UnicodeString(tt.n, []rune(tt.alphabet))

			assert.Len(t, []rune(got), tt.n)
			assertAllowedChars(t, tt.alphabet, got)
		})
	}
}

var dnsLabelTestCases = []struct {
	name       string
	n          int
	panicErrIs error
	panicStr   string
}{
	{
		name:       "n=-128",
		n:          -128,
		panicErrIs: rands.ErrDNSLabelLength,
		panicStr: "rands: DNS labels must be between 1 and 63 characters " +
			"in length",
	},
	{
		name:       "n=0",
		n:          0,
		panicErrIs: rands.ErrDNSLabelLength,
		panicStr: "rands: DNS labels must be between 1 and 63 characters " +
			"in length",
	},
	{name: "n=1", n: 1},
	{name: "n=2", n: 2},
	{name: "n=3", n: 3},
	{name: "n=4", n: 4},
	{name: "n=5", n: 5},
	{name: "n=6", n: 6},
	{name: "n=7", n: 7},
	{name: "n=8", n: 8},
	{name: "n=16", n: 16},
	{name: "n=32", n: 32},
	{name: "n=63", n: 63},
	{
		name:       "n=64",
		n:          64,
		panicErrIs: rands.ErrDNSLabelLength,
		panicStr: "rands: DNS labels must be between 1 and 63 characters " +
			"in length",
	},
	{
		name:       "n=128",
		n:          128,
		panicErrIs: rands.ErrDNSLabelLength,
		panicStr: "rands: DNS labels must be between 1 and 63 characters " +
			"in length",
	},
}

func TestDNSLabel(t *testing.T) {
	t.Parallel()

	for _, tt := range dnsLabelTestCases {
		t.Run(tt.name, func(t *testing.T) {
			// generate lots of labels to increase the chances of catching any
			// obscure bugs
			for i := 0; i < 10000; i++ {
				var got string
				p := recoverPanic(func() {
					got = DNSLabel(tt.n)
				})

				if tt.panicErrIs == nil || tt.panicStr == "" {
					require.Len(t, got, tt.n)
					asserDNSLabel(t, got)
				}

				if tt.panicErrIs != nil {
					require.ErrorIs(t, p.(error), tt.panicErrIs)
				}

				if tt.panicStr != "" {
					require.EqualError(t, p.(error), tt.panicStr)
				}
			}
		})
	}
}

func TestUUID(t *testing.T) {
	t.Parallel()

	m := regexp.MustCompile(
		`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`,
	)

	for i := 0; i < 10000; i++ {
		got := UUID()
		require.Regexp(t, m, got)

		raw := strings.ReplaceAll(got, "-", "")
		b := make([]byte, 16)
		_, err := hex.Decode(b, []byte(raw))
		require.NoError(t, err)

		require.Equal(t, 4, int(b[6]>>4), "version is not 4")
		require.Equal(t, byte(0x80), b[8]&0xc0,
			"variant is not RFC 4122",
		)
	}
}

//
// Helpers
//

var (
	dnsLabelHeadRx = regexp.MustCompile(`^[a-z]$`)
	dnsLabelBodyRx = regexp.MustCompile(`^[a-z0-9-]+$`)
	dnsLabelTailRx = regexp.MustCompile(`^[a-z0-9]$`)
)

func asserDNSLabel(t *testing.T, label string) {
	require.LessOrEqualf(t, len(label), 63,
		`DNS label "%s" is longer than 63 characters`, label,
	)

	require.GreaterOrEqualf(t, len(label), 1,
		`DNS label "%s" is shorter than 1 character`, label,
	)

	if len(label) >= 1 {
		require.Regexpf(t, dnsLabelHeadRx, string(label[0]),
			`DNS label "%s" must start with a-z`, label,
		)
	}
	if len(label) >= 2 {
		require.Regexpf(t, dnsLabelTailRx, string(label[len(label)-1]),
			`DNS label "%s" must end with a-z0-9`, label,
		)
	}
	if len(label) >= 3 {
		require.Regexpf(t, dnsLabelBodyRx, label[1:len(label)-1],
			`DNS label "%s" body must only contain a-z0-9-`, label)
	}
	if len(label) >= 4 {
		require.NotEqualf(t, "--", label[2:4],
			`DNS label "%s" cannot contain "--" as 3rd and 4th char`, label,
		)
	}
}

func assertAllowedChars(t *testing.T, allowed string, s string) {
	invalid := ""
	for _, c := range s {
		if !strings.Contains(allowed, string(c)) &&
			!strings.Contains(invalid, string(c)) {
			invalid += string(c)
		}
	}

	assert.Truef(
		t, len(invalid) == 0, "string contains invalid chars: %s", invalid,
	)
}
