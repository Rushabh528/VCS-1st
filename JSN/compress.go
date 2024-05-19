package JSN

import (
	"bytes"
	"compress/zlib"
)

// I had described Inflate in Session3
// check that code out and its comments
// Try to understand this one by yourself
// Read official docs if needed
func Compress(content []byte) []byte {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(content)
	w.Close()
	return b.Bytes()
}
