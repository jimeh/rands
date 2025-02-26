package uuid

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUUIDv7(t *testing.T) {
	t.Parallel()

	m := regexp.MustCompile(
		`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`,
	)

	// Store timestamps to verify they're increasing.
	var lastTimestampBytes int64
	var lastUUID string

	seen := make(map[string]struct{})
	for i := 0; i < 10000; i++ {
		got, err := NewV7()
		require.NoError(t, err)
		require.Regexp(t, m, got.String())

		if _, ok := seen[got.String()]; ok {
			require.FailNow(t, "duplicate UUID")
		}

		seen[got.String()] = struct{}{}

		// Check version is 7.
		require.Equal(t, 7, int(got[6]>>4), "version is not 7")

		// Check variant is RFC 4122.
		require.Equal(t, byte(0x80), got[8]&0xc0, "variant is not RFC 4122")

		// Extract timestamp bytes.
		timestampBytes := int64(got[0])<<40 | int64(got[1])<<32 |
			int64(got[2])<<24 | int64(got[3])<<16 | int64(got[4])<<8 |
			int64(got[5])

		// Verify timestamp is within 100 milliseconds of current time.
		tsTime := time.UnixMilli(timestampBytes)
		require.WithinDuration(t, time.Now(), tsTime, 100*time.Millisecond,
			"timestamp is not within 100 milliseconds of current time")

		// After the first UUID, verify that UUIDs are monotonically increasing
		if i > 0 && timestampBytes < lastTimestampBytes {
			require.FailNow(t, "UUIDs are not monotonically increasing",
				"current: %s (ts: %d), previous: %s (ts: %d)",
				got, timestampBytes, lastUUID, lastTimestampBytes)
		}

		lastTimestampBytes = timestampBytes
		lastUUID = got.String()
	}
}

func BenchmarkNewV7(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = NewV7()
	}
}

func TestV7Time(t *testing.T) {
	t.Parallel()

	// Define a reference time for testing.
	refTime := time.Date(2023, 5, 15, 12, 0, 0, 0, time.UTC)
	// The timestamp in milliseconds.
	timestampMillis := refTime.UnixMilli()

	// Create bytes for the timestamp (first 6 bytes of UUID).
	timestampBytes := []byte{
		byte(timestampMillis >> 40),
		byte(timestampMillis >> 32),
		byte(timestampMillis >> 24),
		byte(timestampMillis >> 16),
		byte(timestampMillis >> 8),
		byte(timestampMillis),
	}

	tests := []struct {
		name     string
		uuidStr  string
		wantTime time.Time
		wantOk   bool
	}{
		{
			name: "Version 7 UUID",
			uuidStr: func() string {
				var u UUID
				// Set first 6 bytes to timestamp.
				copy(u[:6], timestampBytes)
				// Set version to 7 (0111 as the high nibble of byte 6).
				u[6] = (u[6] & 0x0F) | 0x70
				// Set variant to RFC 4122 (10xx as the high bits of byte 8).
				u[8] = (u[8] & 0x3F) | 0x80

				return u.String()
			}(),
			wantTime: refTime,
			wantOk:   true,
		},
		{
			name: "Version 4 UUID (not time-based)",
			uuidStr: func() string {
				var u UUID
				// Set first 6 bytes to same timestamp to verify it's ignored.
				copy(u[:6], timestampBytes)
				// Set version to 4 (0100 as the high nibble of byte 6).
				u[6] = (u[6] & 0x0F) | 0x40
				// Set variant to RFC 4122 (10xx as the high bits of byte 8).
				u[8] = (u[8] & 0x3F) | 0x80

				return u.String()
			}(),
			wantTime: time.Time{}, // Zero time for non-V7 UUIDs
			wantOk:   false,
		},
		{
			name:     "Zero UUID",
			uuidStr:  "00000000-0000-0000-0000-000000000000",
			wantTime: time.Time{}, // Zero time for version 0
			wantOk:   false,
		},
		{
			name:     "Invalid UUID string",
			uuidStr:  "not-a-valid-uuid",
			wantTime: time.Time{},
			wantOk:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTime, gotOk := V7Time(tt.uuidStr)

			assert.Equal(t, tt.wantOk, gotOk)

			if tt.wantTime.IsZero() {
				assert.True(t, gotTime.IsZero())
			} else {
				// Compare time at millisecond precision.
				assert.Equal(t, tt.wantTime.UnixMilli(), gotTime.UnixMilli())
			}
		})
	}
}

func BenchmarkV7Time(b *testing.B) {
	u, err := NewV7()
	require.NoError(b, err)

	s := u.String()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = V7Time(s)
	}
}
