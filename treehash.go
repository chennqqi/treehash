package treehash

import (
	"bytes"
	"crypto/sha256"
	"hash"
)

const chunkSize = 1024 * 1024

type digest struct {
	buf *bytes.Buffer
}

type node struct {
	parent *node
	lchild *node
	rchild *node
	sha256 []byte
}

func New() hash.Hash {
	return &digest{
		new(bytes.Buffer),
	}
}

func FromHashes(in [][]byte) []byte {
	nodes := make([]*node, len(in))
	for i, b := range in {
		nodes[i] = &node{
			sha256: b,
		}
	}
	root := reduce(nodes)
	return root.sha256
}

func (d *digest) Size() int { return sha256.Size }

func (d *digest) BlockSize() int { return sha256.BlockSize }

func (d *digest) Reset() {
	d.buf = new(bytes.Buffer)
}

func makeChunks(buffer *bytes.Buffer) []*node {
	sha := sha256.New()
	nodes := []*node{}
	chunk := make([]byte, chunkSize)
	for {
		read, err := buffer.Read(chunk)
		if err != nil {
			break
		}
		sha.Reset()
		sha.Write(chunk[:read])
		nodes = append(nodes, &node{sha256: sha.Sum(nil)})

	}
	return nodes
}

func reduce(nodes []*node) *node {
	if len(nodes) == 0 {
		return nil
	}
	if len(nodes) == 1 {
		return nodes[0]
	}
	sha := sha256.New()
	var parentNodes []*node
	for i := 0; i < len(nodes); i += 2 {
		if (i == len(nodes)-1) && (len(nodes)%2 == 1) {
			parent := &node{
				sha256: nodes[i].sha256,
				lchild: nodes[i],
			}
			nodes[i].parent = parent
			parentNodes = append(parentNodes, parent)
		} else {
			sha.Reset()
			sha.Write(nodes[i].sha256)
			sha.Write(nodes[i+1].sha256)
			parent := &node{
				sha256: sha.Sum(nil),
				lchild: nodes[i],
				rchild: nodes[i+1],
			}
			nodes[i].parent, nodes[i+1].parent = parent, parent
			parentNodes = append(parentNodes, parent)
		}
	}
	return reduce(parentNodes)
}

func (d *digest) Sum(in []byte) []byte {
	chunks := makeChunks(d.buf)
	root := reduce(chunks)
	return root.sha256
}

func (d *digest) Write(p []byte) (n int, err error) {
	return d.buf.Write(p)
}
