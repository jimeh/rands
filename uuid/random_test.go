package uuid

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRandom(t *testing.T) {
	t.Parallel()

	m := regexp.MustCompile(
		`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`,
	)

	seen := make(map[string]struct{})

	for i := 0; i < 10000; i++ {
		got, err := NewRandom()
		require.NoError(t, err)
		require.Regexp(t, m, got.String())

		if _, ok := seen[got.String()]; ok {
			require.FailNow(t, "duplicate UUID")
		}

		seen[got.String()] = struct{}{}

		require.Equal(t, 4, int(got[6]>>4), "version is not 4")
		require.Equal(t, byte(0x80), got[8]&0xc0, "variant is not RFC 4122")
	}
}

func BenchmarkNewRandom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = NewRandom()
	}
}
