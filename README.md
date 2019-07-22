# scrypt_playground

## scrypt benchmarking

```
go build
go test -bench=Basic -benchtime=20s

```

## bench results (2.9ghz intel mbp)

### n=512 (1<<9), r=1, p=1, key-len=1, key=64 bytes
- 10 khps

