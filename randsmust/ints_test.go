package randsmust

import (
	"testing"

	"github.com/jimeh/rands"
	"github.com/stretchr/testify/assert"
)

var testIntCases = []struct {
	name       string
	max        int
	panicErrIs error
	panicStr   string
}{
	{
		name:       "n=-2394345",
		max:        -2394345,
		panicErrIs: rands.ErrInvalidMaxInt,
		panicStr:   "rands: max cannot be less than 1",
	},
	{
		name:       "n=-409600",
		max:        -409600,
		panicErrIs: rands.ErrInvalidMaxInt,
		panicStr:   "rands: max cannot be less than 1",
	},
	{
		name:       "n=-1024",
		max:        -1024,
		panicErrIs: rands.ErrInvalidMaxInt,
		panicStr:   "rands: max cannot be less than 1",
	},
	{
		name:       "n=-128",
		max:        -128,
		panicErrIs: rands.ErrInvalidMaxInt,
		panicStr:   "rands: max cannot be less than 1",
	},
	{
		name:       "n=-32",
		max:        -32,
		panicErrIs: rands.ErrInvalidMaxInt,
		panicStr:   "rands: max cannot be less than 1",
	},
	{
		name:       "n=-16",
		max:        -16,
		panicErrIs: rands.ErrInvalidMaxInt,
		panicStr:   "rands: max cannot be less than 1",
	},
	{
		name:       "n=-8",
		max:        -8,
		panicErrIs: rands.ErrInvalidMaxInt,
		panicStr:   "rands: max cannot be less than 1",
	},
	{
		name:       "n=-7",
		max:        -7,
		panicErrIs: rands.ErrInvalidMaxInt,
		panicStr:   "rands: max cannot be less than 1",
	},
	{
		name:       "n=-2",
		max:        -2,
		panicErrIs: rands.ErrInvalidMaxInt,
		panicStr:   "rands: max cannot be less than 1",
	},
	{
		name:       "n=-1",
		max:        -1,
		panicErrIs: rands.ErrInvalidMaxInt,
		panicStr:   "rands: max cannot be less than 1",
	},
	{
		name:       "n=0",
		max:        0,
		panicErrIs: rands.ErrInvalidMaxInt,
		panicStr:   "rands: max cannot be less than 1",
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
	t.Parallel()

	for _, tt := range testIntCases {
		t.Run(tt.name, func(t *testing.T) {
			var got int
			p := recoverPanic(func() {
				got = Int(tt.max)
			})

			if tt.panicErrIs == nil || tt.panicStr == "" {
				assert.GreaterOrEqual(t, got, 0)
				assert.LessOrEqual(t, got, tt.max)
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

func TestInt64(t *testing.T) {
	t.Parallel()

	for _, tt := range testIntCases {
		t.Run(tt.name, func(t *testing.T) {
			var got int64
			p := recoverPanic(func() {
				got = Int64(int64(tt.max))
			})

			if tt.panicErrIs == nil || tt.panicStr == "" {
				assert.GreaterOrEqual(t, got, int64(0))
				assert.LessOrEqual(t, got, int64(tt.max))
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
