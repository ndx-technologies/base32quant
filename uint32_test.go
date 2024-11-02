package base32quant_test

import (
	"fmt"

	"github.com/ndx-technologies/base32quant"
)

func ExampleNewUInt32() {
	u := base32quant.NewUInt32(42)
	fmt.Println(u.UInt32())
	// Output: 42
}

func ExampleUInt32_MarshalText() {
	u := base32quant.NewUInt32(42)
	b, _ := u.MarshalText()
	fmt.Println(string(b))
	// Output: aaaaabk
}

func ExampleUInt32_UnmarshalText() {
	var u base32quant.UInt32
	u.UnmarshalText([]byte("aaaaabk"))
	fmt.Println(u.UInt32())
	// Output: 42
}

func ExampleUInt32_UnmarshalText_error() {
	var u base32quant.UInt32
	fmt.Println(u.UnmarshalText([]byte("aaaaaaaaaabk")))
	// Output: invalid length
}

func ExampleUInt32_String() {
	u := base32quant.NewUInt32(42)
	fmt.Println(u.String())
	// Output: aaaaabk
}

func ExampleNewUInt32FromString() {
	u, _ := base32quant.NewUInt32FromString("aaaaabk")
	fmt.Println(u.UInt32())
	// Output: 42
}
