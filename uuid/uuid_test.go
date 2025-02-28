package uuid

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUUID_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		uuid     UUID
		expected string
	}{
		{
			name:     "Zero UUID",
			uuid:     UUID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expected: "00000000-0000-0000-0000-000000000000",
		},
		{
			name: "Random UUID",
			uuid: UUID{
				0x12, 0x34, 0x56, 0x78,
				0x9a, 0xbc,
				0xde, 0xf0,
				0x12, 0x34,
				0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
			},
			expected: "12345678-9abc-def0-1234-56789abcdef0",
		},
		{
			name: "UUID with max values",
			uuid: UUID{
				0xff, 0xff, 0xff, 0xff,
				0xff, 0xff,
				0xff, 0xff,
				0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			},
			expected: "ffffffff-ffff-ffff-ffff-ffffffffffff",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.uuid.String()
			assert.Equal(t, tt.expected, got)
		})
	}
}

func BenchmarkUUID_String(b *testing.B) {
	u := UUID{
		0x12, 0x34, 0x56, 0x78,
		0x9a, 0xbc,
		0xde, 0xf0,
		0x12, 0x34,
		0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = u.String()
	}
}

func TestFromBytes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		bytes   []byte
		want    UUID
		wantErr error
	}{
		{
			name: "Valid 16 bytes",
			bytes: []byte{
				0x12,
				0x34,
				0x56,
				0x78,
				0x9a,
				0xbc,
				0xde,
				0xf0,
				0x12,
				0x34,
				0x56,
				0x78,
				0x9a,
				0xbc,
				0xde,
				0xf0,
			},
			want: UUID{
				0x12, 0x34, 0x56, 0x78,
				0x9a, 0xbc,
				0xde, 0xf0,
				0x12, 0x34,
				0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
			},
			wantErr: nil,
		},
		{
			name:    "Empty bytes",
			bytes:   []byte{},
			want:    UUID{},
			wantErr: ErrInvalidLength,
		},
		{
			name:    "Too few bytes",
			bytes:   []byte{0x12, 0x34, 0x56, 0x78, 0x9a},
			want:    UUID{},
			wantErr: ErrInvalidLength,
		},
		{
			name: "Too many bytes",
			bytes: []byte{
				0x12,
				0x34,
				0x56,
				0x78,
				0x9a,
				0xbc,
				0xde,
				0xf0,
				0x12,
				0x34,
				0x56,
				0x78,
				0x9a,
				0xbc,
				0xde,
				0xf0,
				0xab,
			},
			want:    UUID{},
			wantErr: ErrInvalidLength,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromBytes(tt.bytes)

			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkFromBytes(b *testing.B) {
	bytes := []byte{
		0x12, 0x34, 0x56, 0x78,
		0x9a, 0xbc,
		0xde, 0xf0,
		0x12, 0x34,
		0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FromBytes(bytes)
	}
}

func TestFromString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		str     string
		want    UUID
		wantErr error
	}{
		{
			name: "Valid UUID string",
			str:  "12345678-9abc-def0-1234-56789abcdef0",
			want: UUID{
				0x12, 0x34, 0x56, 0x78,
				0x9a, 0xbc,
				0xde, 0xf0,
				0x12, 0x34,
				0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
			},
			wantErr: nil,
		},
		{
			name:    "Empty string",
			str:     "",
			want:    UUID{},
			wantErr: ErrInvalidLength,
		},
		{
			name:    "Too short string",
			str:     "12345678-9abc-def0-1234-56789abcde",
			want:    UUID{},
			wantErr: ErrInvalidLength,
		},
		{
			name:    "Too long string",
			str:     "12345678-9abc-def0-1234-56789abcdef0a",
			want:    UUID{},
			wantErr: ErrInvalidLength,
		},
		{
			name:    "Invalid characters",
			str:     "12345678-9abc-defg-1234-56789abcdef0",
			want:    UUID{},
			wantErr: errors.New("encoding/hex: invalid byte: U+0067 'g'"),
		},
		{
			name:    "Zero UUID",
			str:     "00000000-0000-0000-0000-000000000000",
			want:    UUID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			wantErr: nil,
		},
		{
			name: "Max value UUID",
			str:  "ffffffff-ffff-ffff-ffff-ffffffffffff",
			want: UUID{
				0xff, 0xff, 0xff, 0xff,
				0xff, 0xff,
				0xff, 0xff,
				0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromString(tt.str)

			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkFromString(b *testing.B) {
	uuidStr := "12345678-9abc-def0-1234-56789abcdef0"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FromString(uuidStr)
	}
}

func TestUUID_Time(t *testing.T) {
	t.Parallel()

	// Define a reference time for testing.
	refTime := time.Date(2023, 5, 15, 12, 0, 0, 0, time.UTC)
	// The timestamp in milliseconds (Unix timestamp * 1000 in 6 bytes).
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
		uuid     UUID
		wantTime time.Time
		wantOk   bool
	}{
		{
			name: "Version 7 UUID",
			uuid: func() UUID {
				var u UUID
				// Set first 6 bytes to timestamp.
				copy(u[:6], timestampBytes)
				// Set version to 7 (0111 as the high nibble of byte 6).
				u[6] = (u[6] & 0x0F) | 0x70
				// Set variant to RFC 4122 (10xx as the high bits of byte 8).
				u[8] = (u[8] & 0x3F) | 0x80

				return u
			}(),
			wantTime: refTime,
			wantOk:   true,
		},
		{
			name: "Version 4 UUID (not time-based)",
			uuid: func() UUID {
				var u UUID
				// Set first 6 bytes to same timestamp to verify it's ignored.
				copy(u[:6], timestampBytes)
				// Set version to 4 (0100 as the high nibble of byte 6).
				u[6] = (u[6] & 0x0F) | 0x40

				return u
			}(),
			wantTime: time.Time{}, // Zero time for non-V7 UUIDs
			wantOk:   false,
		},
		{
			name:     "Zero UUID",
			uuid:     UUID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			wantTime: time.Time{}, // Zero time for version 0
			wantOk:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTime, gotOk := tt.uuid.Time()

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

func BenchmarkUUID_Time(b *testing.B) {
	uuid, err := NewV7()
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = uuid.Time()
	}
}

func TestUUID_Version(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		uuid UUID
		want int
	}{
		{
			name: "Version 0 (invalid/nil UUID)",
			uuid: UUID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: 0,
		},
		{
			name: "Version 1 (time-based)",
			uuid: UUID{
				0x12, 0x34, 0x56, 0x78,
				0x9a, 0xbc,
				0x10, 0xf0, // 0x10 = version 1 in top nibble
				0x12, 0x34,
				0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
			},
			want: 1,
		},
		{
			name: "Version 2 (DCE Security)",
			uuid: UUID{
				0x12, 0x34, 0x56, 0x78,
				0x9a, 0xbc,
				0x20, 0xf0, // 0x20 = version 2 in top nibble
				0x12, 0x34,
				0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
			},
			want: 2,
		},
		{
			name: "Version 3 (name-based, MD5)",
			uuid: UUID{
				0x12, 0x34, 0x56, 0x78,
				0x9a, 0xbc,
				0x30, 0xf0, // 0x30 = version 3 in top nibble
				0x12, 0x34,
				0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
			},
			want: 3,
		},
		{
			name: "Version 4 (random)",
			uuid: UUID{
				0x12, 0x34, 0x56, 0x78,
				0x9a, 0xbc,
				0x40, 0xf0, // 0x40 = version 4 in top nibble
				0x12, 0x34,
				0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
			},
			want: 4,
		},
		{
			name: "Version 5 (name-based, SHA-1)",
			uuid: UUID{
				0x12, 0x34, 0x56, 0x78,
				0x9a, 0xbc,
				0x50, 0xf0, // 0x50 = version 5 in top nibble
				0x12, 0x34,
				0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
			},
			want: 5,
		},
		{
			name: "Version 7 (time-ordered)",
			uuid: UUID{
				0x12, 0x34, 0x56, 0x78,
				0x9a, 0xbc,
				0x70, 0xf0, // 0x70 = version 7 in top nibble
				0x12, 0x34,
				0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
			},
			want: 7,
		},
		{
			name: "Version 8 (custom)",
			uuid: UUID{
				0x12, 0x34, 0x56, 0x78,
				0x9a, 0xbc,
				0x80, 0xf0, // 0x80 = version 8 in top nibble
				0x12, 0x34,
				0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
			},
			want: 8,
		},
		{
			name: "Version 15 (theoretical max)",
			uuid: UUID{
				0x12, 0x34, 0x56, 0x78,
				0x9a, 0xbc,
				0xf0, 0xf0, // 0xf0 = version 15 in top nibble
				0x12, 0x34,
				0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
			},
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.uuid.Version()

			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkUUID_Version(b *testing.B) {
	uuid, err := NewV7()
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uuid.Version()
	}
}
