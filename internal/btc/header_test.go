package btc


import "testing"


func TestHeaderSerialize_Length(t *testing.T) {
prev := [32]byte{}
merkle := [32]byte{}
h := BlockHeader{
Version: 0x20000000,
PrevBlockLE: prev,
MerkleRootLE: merkle,
Time: 1234567890,
Bits: 0x1f0fffff,
Nonce: 42,
}
ser := h.Serialize()
if len(ser) != 80 {
t.Fatalf("serialized header must be 80 bytes, got %d", len(ser))
}
}


func TestHeaderHash_Deterministic(t *testing.T) {
// Deterministic across runs when inputs are constant
var prevBE, merkleBE [32]byte
for i := 0; i < 32; i++ { prevBE[i] = byte(i); merkleBE[i] = byte(32+i) }
h := BlockHeader{
Version: 0x20000000,
PrevBlockLE: ToLittleEndian32(prevBE),
MerkleRootLE: ToLittleEndian32(merkleBE),
Time: 1,
Bits: 0x1f0fffff,
Nonce: 0,
}
exp := "4ff6280a5aa0cdff30aca6345d0dcf3024d1f9253eae31d48d7360a398e95540"
hash := h.HashBE()
if ToHex(hash[:]) != exp {
t.Fatalf("unexpected header hash: got %s want %s", ToHex(hash[:]), exp)
}
}
