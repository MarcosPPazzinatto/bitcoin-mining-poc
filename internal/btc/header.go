package btc


import (
"bytes"
"encoding/binary"
)


// BlockHeader is the 80-byte Bitcoin header (fields serialized little-endian).
type BlockHeader struct {
Version uint32
PrevBlockLE [32]byte // little-endian
MerkleRootLE [32]byte // little-endian
Time uint32
Bits uint32 // compact target
Nonce uint32
}


func (h *BlockHeader) Serialize() []byte {
buf := bytes.NewBuffer(make([]byte, 0, 80))
_ = binary.Write(buf, binary.LittleEndian, h.Version)
buf.Write(h.PrevBlockLE[:])
buf.Write(h.MerkleRootLE[:])
_ = binary.Write(buf, binary.LittleEndian, h.Time)
_ = binary.Write(buf, binary.LittleEndian, h.Bits)
_ = binary.Write(buf, binary.LittleEndian, h.Nonce)
return buf.Bytes()
}


// HashBE returns the double SHA-256 of the serialized header in big-endian.
func (h *BlockHeader) HashBE() [32]byte {
return DoubleSha256(h.Serialize())
}
