package internal

import (
	"fmt"
)

func VariablesAndDeclaration() {
	// Explicit declaration with initialization
	var a = "initial"
	fmt.Println(a)

	// Declaration of multiple variables of the same type
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Declaration with type inference
	var d = true
	fmt.Println(d)

	// Declaration without initialization, zero value
	var e int
	fmt.Println(e)

	// Short declaration with type inference
	f := "apple"
	fmt.Println(f)

	// Explicit type declaration
	var g float64 = 3.14
	fmt.Println(g)

	var h uint8 = 255
	fmt.Println(h)

	// Boolean variable declaration
	var i bool = false
	fmt.Println(i)

	// Complex number variable declaration
	var j complex128 = complex(1, 2)
	fmt.Println(j)

	// Byte variable declaration (alias for uint8)
	var k byte = 100
	fmt.Println(k)

	// Rune variable declaration (alias for int32)
	var l rune = 'G'
	fmt.Println(l)

	// Declaration of multiple variables of different types using parentheses
	var (
		m int    = 42
		n string = "hello"
		o bool   = true
	)
	fmt.Println(m, n, o)
}
