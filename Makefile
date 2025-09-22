.PHONY: build run-sim test tidy


build:
go build ./...


run-sim:
go run ./cmd/powsim --max-nonces 50000000 --bits 1f0fffff --progress-every 2000000


test:
go test ./...


tidy:
go mod tidy
