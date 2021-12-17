package randsmust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	t.Parallel()

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Bytes(tt.n)

			assert.Len(t, got, tt.n)
		})
	}
}
