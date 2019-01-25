package treehash

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestTreeHash(t *testing.T) {
	// create 1.5MB empty file
	b := make([]byte, 1536000)
	th := New()
	io.Copy(th, bytes.NewReader(b))
	hash := fmt.Sprintf("%x", th.Sum(nil))
	if hash != "ca1f50f208ffe93e48045d48359cf31109b852601016c33db554735964d9e636" {
		t.Fatal("incorrect checksum")
	}
}
