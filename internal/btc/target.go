package btc


import (
"encoding/binary"
"errors"
"math/big"
)


// TargetFromBits decodes the compact target (bits) to a big integer target.
// bits: 1-byte exponent (E) + 3-byte mantissa (M)
// target = M * 256^(E-3)
func TargetFromBits(bits uint32) (*big.Int, error) {
exp := int(bits >> 24)
mant := bits & 0x007fffff
if bits&0x00800000 != 0 {
return nil, errors.New("invalid compact bits (negative)")
}
M := new(big.Int).SetUint64(uint64(mant))
shift := 8 * (exp - 3)
if shift < 0 {
return new(big.Int).Rsh(M, uint(-shift)), nil
}
return new(big.Int).Lsh(M, uint(shift)), nil
}


// HashMeetsTarget returns true if the big-endian hash <= target.
func HashMeetsTarget(hashBE [32]byte, target *big.Int) bool {
h := new(big.Int).SetBytes(hashBE[:])
return h.Cmp(target) <= 0
}


// ToLittleEndian32 flips a 32-byte big-endian array to little-endian order.
func ToLittleEndian32(bytesBE [32]byte) (out [32]byte) {
for i := 0; i < 32; i++ {
out[i] = bytesBE[31-i]
}
return
}


// U32LE converts a uint32 to 4 little-endian bytes.
func U32LE(v uint32) [4]byte {
var b [4]byte
binary.LittleEndian.PutUint32(b[:], v)
return b
}
