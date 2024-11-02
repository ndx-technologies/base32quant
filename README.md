fast `base32` encoding into `uint32` and `uint64`

```bash
$ go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: github.com/ndx-technologies/mm-go/base32
cpu: Apple M3 Max
BenchmarkEncodeDecode_Standard/uint32-16                32816378                36.46 ns/op           16 B/op          2 allocs/op
BenchmarkEncodeDecode_Standard/uint64-16                23218821                50.55 ns/op           24 B/op          2 allocs/op
BenchmarkEncodeDecode/uint32-16                         128822170                9.676 ns/op           0 B/op          0 allocs/op
BenchmarkEncodeDecode/uint64-16                         93620481                12.96 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/ndx-technologies/mm-go/base32        7.052s
```
