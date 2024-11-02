package base32quant_test

import (
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"testing"

	base32quant "github.com/ndx-technologies/base32quant"
)

func Benchmark__________________________________(b *testing.B) { b.SkipNow() }

func BenchmarkEncodeDecode_Standard(b *testing.B) {
	enc := base32.StdEncoding.WithPadding(base32.NoPadding)

	b.Run("uint32", func(b *testing.B) {
		v := make([]byte, 7)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := range v {
				v[j] = 0
			}

			enc.Encode(v, binary.BigEndian.AppendUint32(nil, uint32(i)))

			g := [4]byte{}
			_, err := enc.Decode(g[:], v)

			o := binary.BigEndian.Uint32(g[:])

			if o != uint32(i) || err != nil {
				b.Fatalf("invalid value: %d %d %v, err: %s", o, i, v, err)
			}
		}
	})

	b.Run("uint64", func(b *testing.B) {
		v := make([]byte, 13)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := range v {
				v[j] = 0
			}

			enc.Encode(v, binary.BigEndian.AppendUint64(nil, uint64(i)))

			g := [8]byte{}
			_, err := enc.Decode(g[:], v)

			o := binary.BigEndian.Uint64(g[:])

			if o != uint64(i) || err != nil {
				b.Fatalf("invalid value: %d %d %v, err: %s", o, i, v, err)
			}
		}
	})
}

func BenchmarkEncodeDecode(b *testing.B) {
	b.Run("uint32", func(b *testing.B) {
		v := make([]byte, 7)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := range v {
				v[j] = 0
			}

			base32quant.Encode(v, uint32(i))
			o, err := base32quant.Decode[uint32](v)

			if o != uint32(i) || err != nil {
				b.Fatalf("invalid value: %d %d %v, err: %s", o, i, v, err)
			}
		}
	})

	b.Run("uint64", func(b *testing.B) {
		v := make([]byte, 13)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := range v {
				v[j] = 0
			}

			base32quant.Encode(v, uint64(i))
			o, err := base32quant.Decode[uint64](v)

			if o != uint64(i) || err != nil {
				b.Fatalf("invalid value: %d %d %v, err: %s", o, i, v, err)
			}
		}
	})
}

func ExampleEncode() {
	var x uint32 = 1234567890

	b := make([]byte, 7)
	base32quant.Encode(b, x)

	fmt.Println(string(b))
	// Output: bezmaws
}

func ExampleDecode() {
	x, _ := base32quant.Decode[uint32]([]byte("bezmaws"))
	fmt.Println(x)
	// Output: 1234567890
}

func ExampleDecode_mixedCase() {
	x, _ := base32quant.Decode[uint32]([]byte("beZmAwS"))
	fmt.Println(x)
	// Output: 1234567890
}

func TestEncodeDecode(t *testing.T) {
	s := "-1"
	_, err := base32quant.Decode[uint32]([]byte(s))
	if err == nil {
		t.Fatalf("must be error")
	}
}

func FuzzEncodeDecode_32(f *testing.F) {
	b := make([]byte, 7)

	f.Add(uint32(0))
	f.Add(uint32(11))
	f.Add(uint32(32123))
	f.Add(uint32(222))

	f.Fuzz(func(t *testing.T, v uint32) {
		base32quant.Encode(b, v)
		o, err := base32quant.Decode[uint32](b)

		if o != v || err != nil {
			t.Fatalf("invalid value: %d %d %v, err: %s", o, v, b, err)
		}
	})
}

func FuzzEncodeDecode_64(f *testing.F) {
	b := make([]byte, 13)

	f.Add(uint64(0))
	f.Add(uint64(11))

	f.Fuzz(func(t *testing.T, v uint64) {
		base32quant.Encode(b, v)
		o, err := base32quant.Decode[uint64](b)

		if o != v || err != nil {
			t.Fatalf("invalid value: %d %d %v, err: %s", o, v, b, err)
		}
	})
}
