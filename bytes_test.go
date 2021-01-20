package rands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Bytes(tt.n)

			assert.Len(t, got, tt.n)
		})
	}
}

func BenchmarkBytes(b *testing.B) {
	for _, tt := range testCases {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, _ = Bytes(tt.n)
			}
		})
	}
}
