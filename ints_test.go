package rands

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testIntCases = []struct {
	name   string
	max    int
	errIs  error
	errStr string
}{
	{
		name:   "n=-2394345",
		max:    -2394345,
		errIs:  errInvalidMaxInt,
		errStr: "rands: max cannot be less than 1",
	},
	{
		name:   "n=-409600",
		max:    -409600,
		errIs:  errInvalidMaxInt,
		errStr: "rands: max cannot be less than 1",
	},
	{
		name:   "n=-1024",
		max:    -1024,
		errIs:  errInvalidMaxInt,
		errStr: "rands: max cannot be less than 1",
	},
	{
		name:   "n=-128",
		max:    -128,
		errIs:  errInvalidMaxInt,
		errStr: "rands: max cannot be less than 1",
	},
	{
		name:   "n=-32",
		max:    -32,
		errIs:  errInvalidMaxInt,
		errStr: "rands: max cannot be less than 1",
	},
	{
		name:   "n=-16",
		max:    -16,
		errIs:  errInvalidMaxInt,
		errStr: "rands: max cannot be less than 1",
	},
	{
		name:   "n=-8",
		max:    -8,
		errIs:  errInvalidMaxInt,
		errStr: "rands: max cannot be less than 1",
	},
	{
		name:   "n=-7",
		max:    -7,
		errIs:  errInvalidMaxInt,
		errStr: "rands: max cannot be less than 1",
	},
	{
		name:   "n=-2",
		max:    -2,
		errIs:  errInvalidMaxInt,
		errStr: "rands: max cannot be less than 1",
	},
	{
		name:   "n=-1",
		max:    -1,
		errIs:  errInvalidMaxInt,
		errStr: "rands: max cannot be less than 1",
	},
	{
		name:   "n=0",
		max:    0,
		errIs:  errInvalidMaxInt,
		errStr: "rands: max cannot be less than 1",
	},
	{name: "n=1", max: 1},
	{name: "n=2", max: 2},
	{name: "n=7", max: 7},
	{name: "n=8", max: 8},
	{name: "n=16", max: 16},
	{name: "n=32", max: 32},
	{name: "n=128", max: 128},
	{name: "n=1024", max: 1024},
	{name: "n=409600", max: 409600},
	{name: "n=2394345", max: 2394345},
}

func TestInt(t *testing.T) {
	for _, tt := range testIntCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int(tt.max)

			if tt.errIs == nil || tt.errStr == "" {
				assert.GreaterOrEqual(t, got, 0)
				assert.LessOrEqual(t, got, tt.max)
			}

			if tt.errIs != nil {
				assert.True(t, errors.Is(err, errInvalidMaxInt))
			}

			if tt.errStr != "" {
				assert.EqualError(t, err, tt.errStr)
			}
		})
	}
}

func BenchmarkInt(b *testing.B) {
	for _, tt := range testIntCases {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, _ = Int(tt.max)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	for _, tt := range testIntCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64(int64(tt.max))

			if tt.errIs == nil || tt.errStr == "" {
				assert.GreaterOrEqual(t, got, int64(0))
				assert.LessOrEqual(t, got, int64(tt.max))
			}

			if tt.errIs != nil {
				assert.True(t, errors.Is(err, errInvalidMaxInt))
			}

			if tt.errStr != "" {
				assert.EqualError(t, err, tt.errStr)
			}
		})
	}
}

func BenchmarkInt64(b *testing.B) {
	for _, tt := range testIntCases {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, _ = Int64(int64(tt.max))
			}
		})
	}
}
