package btc

import (
	"crypto/sha256"
)

func Sha256(b []byte) [32]byte {
	return sha256.Sum256(b)
}

func DoubleSha256(b []byte) [32]byte {
	h1 := sha256.Sum256(b)
	return sha256.Sum256(h1[:])
}

func ToHex(b []byte) string {
	const hexdigits = "0123456789abcdef"
	out := make([]byte, len(b)*2)
	for i, v := range b {
		out[i*2] = hexdigits[v>>4]
		out[i*2+1] = hexdigits[v&0x0f]
	}
	return string(out)
}

