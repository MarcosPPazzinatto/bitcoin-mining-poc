# Bitcoin Mining PoC (Go)

This repository provides an educational proof-of-concept demonstrating Bitcoin Proof-of-Work.

## Features
- **PoW Simulator (offline)**: Brute-forces a nonce against a chosen compact target (`bits`).
- **Regtest Miner (stub)**: Planned extension to connect with `bitcoind` in regtest mode using RPC.

> This project is strictly educational. CPU mining on Bitcoin mainnet is not feasible.

## Requirements
- Go 1.22 or higher
- (Optional, for regtest) Bitcoin Core 27+ (`bitcoind`, `bitcoin-cli`)

## Build
```
go build ./...
```
## Run PoW Simulator
```
go run ./cmd/powsim --max-nonces 50000000 --bits 1f0fffff --progress-every 2000000
```

## Flags

- ``--max-nonces`` Maximum attempts (default: 20000000)
- ``--bits Compact`` target in hex (default: 1f07ffff)
- ``--progress-every`` Print progress after N nonces (default: 4000000)

## Next Steps: Regtest Miner

### 1. Start Bitcoin Core in regtest mode:

```
bitcoind -regtest -daemon
bitcoin-cli -regtest createwallet miner || true
bitcoin-cli -regtest -generate 101
```

### 2. Implement RPC client for:


- getblocktemplate
- block assembly (coinbase + merkle root)
- nonce loop and submitblock





















