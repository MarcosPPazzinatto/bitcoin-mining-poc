package btc


import (
"encoding/hex"
"testing"
)


func TestDoubleSha256_KnownVectors(t *testing.T) {
cases := []struct{
name string
in []byte
expHex string
}{
{"empty", []byte(""), "5df6e0e2761359d30a8275058e299fcc0381534545f55cf43e41983f5d4c9456"},
{"abc", []byte("abc"), "4f8b42c22dd3729b519ba6f68d2da7cc5b2d606d05daed5ad5128cc03e6c6358"},
}
for _, tc := range cases {
got := DoubleSha256(tc.in)
if hex.EncodeToString(got[:]) != tc.expHex {
t.Fatalf("%s: unexpected double SHA256: got %s want %s", tc.name, hex.EncodeToString(got[:]), tc.expHex)
}
}
}


func TestDoubleSha256_AllZero80(t *testing.T) {
in := make([]byte, 80)
exp := "4be7570e8f70eb093640c8468274ba759745a7aa2b7d25ab1e0421b259845014"
got := DoubleSha256(in)
if hex.EncodeToString(got[:]) != exp {
t.Fatalf("unexpected double SHA256: got %s want %s", hex.EncodeToString(got[:]), exp)
}
}
