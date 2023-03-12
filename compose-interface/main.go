package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

type HashReader interface {
	io.Reader
	hash() string
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	hash := sha1.Sum(h.buf.Bytes())
	return hex.EncodeToString(hash[:])
}

func hashAndBroadcast(r HashReader) error {
	hash := r.hash()
	fmt.Println(hash)

	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("string of the bytes: ", string(b))

	return nil
}

func main() {
	payload := []byte("Composing..")
	hashAndBroadcast(NewHashReader(payload))

}
