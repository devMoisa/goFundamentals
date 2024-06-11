package internal

import "fmt"

func DataTypes() {
	// Boolean
	var a bool = true

	// Integers
	var b int = 5
	var c int8 = 127
	var d int16 = 32767
	var e int32 = 2147483647
	var f int64 = 9223372036854775807

	// Unsigned Integers
	var g uint = 5
	var h uint8 = 255
	var i uint16 = 65535
	var j uint32 = 4294967295
	var k uint64 = 18446744073709551615

	// Float
	var l float32 = 3.14
	var m float64 = 3.141592653589793

	// Complex numbers
	var n complex64 = 1 + 2i
	var o complex128 = 1 + 2i

	// Byte (alias for uint8)
	var p byte = 255

	// Rune (alias for int32)
	var q rune = 'A'

	// Unsigned integer pointer
	var r uintptr = 123456

	// String
	var s string = "Hello Fri3nd"

	fmt.Println("Boolean: ", a)

	fmt.Println("Integer (int): ", b)
	fmt.Println("Integer (int8): ", c)
	fmt.Println("Integer (int16): ", d)
	fmt.Println("Integer (int32): ", e)
	fmt.Println("Integer (int64): ", f)

	fmt.Println("Unsigned Integer (uint): ", g)
	fmt.Println("Unsigned Integer (uint8): ", h)
	fmt.Println("Unsigned Integer (uint16): ", i)
	fmt.Println("Unsigned Integer (uint32): ", j)
	fmt.Println("Unsigned Integer (uint64): ", k)

	fmt.Println("Float (float32): ", l)
	fmt.Println("Float (float64): ", m)

	fmt.Println("Complex (complex64): ", n)
	fmt.Println("Complex (complex128): ", o)

	fmt.Println("Byte: ", p)
	fmt.Println("Rune: ", q)

	fmt.Println("Unsigned Integer Pointer: ", r)

	fmt.Println("String: ", s)
}
