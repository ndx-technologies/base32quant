package base32quant_test

import (
	"fmt"

	"github.com/ndx-technologies/base32quant"
)

func ExampleNewUInt64() {
	u := base32quant.NewUInt64(42)
	fmt.Println(u.UInt64())
	// Output: 42
}

func ExampleUInt64_MarshalText() {
	u := base32quant.NewUInt64(42)
	b, _ := u.MarshalText()
	fmt.Println(string(b))
	// Output: aaaaaaaaaaabk
}

func ExampleUInt64_UnmarshalText() {
	var u base32quant.UInt64
	u.UnmarshalText([]byte("aaaaaaaaaaabk"))
	fmt.Println(u.UInt64())
	// Output: 42
}

func ExampleUInt64_UnmarshalText_error() {
	var u base32quant.UInt64
	fmt.Println(u.UnmarshalText([]byte("aaaaaaaaaabk")))
	// Output: invalid length
}

func ExampleUInt64_String() {
	u := base32quant.NewUInt64(42)
	fmt.Println(u.String())
	// Output: aaaaaaaaaaabk
}

func ExampleNewUInt64FromString() {
	u, _ := base32quant.NewUInt64FromString("aaaaaaaaaaabk")
	fmt.Println(u.UInt64())
	// Output: 42
}
