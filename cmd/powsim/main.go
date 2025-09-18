package main
func random32() [32]byte {
var b [32]byte
_, _ = rand.Read(b[:])
return b
}


func main() {
var (
maxNonces uint64
bitsStr string
progressN uint64
)


flag.Uint64Var(&maxNonces, "max-nonces", 20_000_000, "max attempts")
flag.StringVar(&bitsStr, "bits", "1f07ffff", "compact target in hex (e.g., 1f0fffff)")
flag.Uint64Var(&progressN, "progress-every", 4_000_000, "print progress every N nonces")
flag.Parse()


bitsVal, err := hex.DecodeString(bitsStr)
if err != nil || len(bitsVal) > 4 {
panic("invalid bits hex")
}
bits := uint32(0)
for _, v := range bitsVal {
bits = (bits << 8) | uint32(v)
}


prevBE := random32()
merkleBE := random32()


header := btc.BlockHeader{
Version: 0x20000000,
PrevBlockLE: btc.ToLittleEndian32(prevBE),
MerkleRootLE: btc.ToLittleEndian32(merkleBE),
Time: uint32(time.Now().Unix()),
Bits: bits,
Nonce: 0,
}


target, err := btc.TargetFromBits(bits)
if err != nil {
panic(err)
}


start := time.Now()
found := false


for n := uint64(0); n < maxNonces; n++ {
header.Nonce = uint32(n)
h := header.HashBE()
if btc.HashMeetsTarget(h, target) {
fmt.Printf("Found nonce: %d\n", n)
fmt.Printf("Header hash (BE): %s\n", btc.ToHex(h[:]))
fmt.Printf("Elapsed: %v\n", time.Since(start))
found = true
break
}
if progressN > 0 && (n%progressN) == 0 {
el := time.Since(start).Seconds()
rate := float64(n+1) / el
fmt.Printf("... tried nonce=%d (%.0f nonces/s)\n", n, rate)
}
}


if !found {
fmt.Printf("No valid nonce within %d attempts. Elapsed: %v\n", maxNonces, time.Since(start))
fmt.Println("Tip: increase target (lower difficulty), e.g. --bits 1f0fffff")
}
}
