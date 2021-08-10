package main

import (
	"fmt"
)

func main() {
	// declaring a boolean variable
	var a bool = false
	fmt.Printf("%-10T %-3v \n", a, a)

	// default value for boolean
	var b bool
	fmt.Printf("%-10T %-3v \n", b, b)

	// assigning boolean with equations
	c := 1 == 1
	fmt.Printf("%-10T %-3v \n", c, c)

	// declare unsinged integer unit8 goes from 0-255 (2^8 - 1)
	// unit16 goes from 0-65535 (2^16 - 1) unit32 goes from 0-2^32 - 1
	// no unit64 but can use a byte
	var d uint16 = 42
	fmt.Printf("%-10T %-3v \n", d, d)

	fmt.Println()

	e := 10
	f := 3
	fmt.Println("e+f = ", e+f)
	fmt.Println("e-f = ", e-f)
	fmt.Println("e*f = ", e*f)
	fmt.Println("e/f = ", e/f)
	fmt.Println("e%f = ", e%f)
	fmt.Println()

	var g int = 10
	var h int8 = 3
	// need to explicit cast into the same type before addition or it's gonna give an error
	fmt.Println("f+g= ", g+int(h))
	fmt.Println()

	// e in binary 1010
	// f in binary 0011

	fmt.Println("e&f = ", e&f)   // AND 		0010 in binary = 2
	fmt.Println("e|f = ", e|f)   // OR  		1011 in binary = 11
	fmt.Println("e^f = ", e^f)   // NOT			1001 in binary = 9
	fmt.Println("e&^f = ", e&^f) // ANDNOT 		1000 in binary = 8
	fmt.Println()

	var i int = 8
	fmt.Println("i= ", i)
	fmt.Println("i<<3 = ", i<<3) // 1000 shift 3 to the left 0100 0000 = 64 = 2^3 * 2^3
	fmt.Println("i>>3 = ", i>>3) // 1000 shift 3 to the right 0001  = 1 = 2^3 / 2^3
}
