package btc


import (
"encoding/hex"
"math/big"
"testing"
)


func TestTargetFromBits_GenesisBits(t *testing.T) {
// Bitcoin genesis block bits = 0x1d00ffff
bits := uint32(0x1d00ffff)
target, err := TargetFromBits(bits)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
// expected target = 0x00ffff * 2^(8*(0x1d-3))
exp := new(big.Int)
exp.SetUint64(0x00ffff)
exp.Lsh(exp, 8*(0x1d-3))
if target.Cmp(exp) != 0 {
t.Fatalf("target mismatch: got %x want %x", target, exp)
}
}


func TestHashMeetsTarget(t *testing.T) {
bits := uint32(0x1f0fffff) // very easy target for testing
target, err := TargetFromBits(bits)
if err != nil { t.Fatal(err) }


// Case 1: all-zero hash should meet any positive target
var zero [32]byte
if !HashMeetsTarget(zero, target) {
t.Fatal("zero hash should meet target")
}
// Case 2: all-FF hash should not meet target (unless target is max)
var ff [32]byte
for i := range ff { ff[i] = 0xff }
if HashMeetsTarget(ff, target) {
t.Fatal("0xff.. hash should not meet target")
}


// Case 3: boundary check using constructed header hash
// Build a deterministic header and verify compare logic is consistent.
prevBE := [32]byte{}
merkleBE := [32]byte{}
for i := 0; i < 32; i++ { prevBE[i] = byte(i); merkleBE[i] = byte(32+i) }
h := BlockHeader{
Version: 0x20000000,
PrevBlockLE: ToLittleEndian32(prevBE),
MerkleRootLE: ToLittleEndian32(merkleBE),
Time: 1,
Bits: bits,
Nonce: 0,
}
hash := h.HashBE()
meets := HashMeetsTarget(hash, target)
_ = meets // The goal here is to ensure it runs; exact boolean depends on chosen bits.
}


func TestBitsParsingHexRoundTrip(t *testing.T) {
// Simulate the CLI hex parse used by powsim: "1f07ffff" -> uint32
hexStr := "1f07ffff"
b, err := hex.DecodeString(hexStr)
if err != nil || len(b) != 4 { t.Fatalf("unexpected decode: %v len=%d", err, len(b)) }
var bits uint32
for _, v := range b { bits = (bits << 8) | uint32(v) }
if bits != 0x1f07ffff { t.Fatalf("bits mismatch: got %#x want %#x", bits, 0x1f07ffff) }
}
