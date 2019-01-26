package treehash

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeHash(t *testing.T) {
	tests := []struct {
		desc     string
		numBytes uint
		checksum string
	}{
		{
			desc:     "test for less than 1MB",
			numBytes: 500000,
			checksum: "6bb6aefaeaa4e19112e566b467c4301463a30b0a15b9c8248a00ed9cd8e5946b",
		},
		{
			desc:     "test for even number of file chunks",
			numBytes: 1500000,
			checksum: "f1cdd9e081e997f7a62522dce66d827deac748852a4a133b29458e02b43fc124",
		},
		{
			desc:     "test for odd number of file chunks",
			numBytes: 2500000,
			checksum: "ee1ddc44cb2d853ee2197d3048dfc4fc456dd71d0ad4ec745fab7a50a50ab401",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			b := make([]byte, test.numBytes)
			th := New()
			io.Copy(th, bytes.NewReader(b))
			assert.Equal(t, fmt.Sprintf("%x", th.Sum(nil)), test.checksum)
		})
	}
}
