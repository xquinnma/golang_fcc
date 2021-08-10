package main

import (
	"fmt"
	"strconv"
)

// variable decalaration at package scope
var f int = 5

// uppercase variable is a global variable, G is a global variable
var G int = 7

func main() {
	// ways of declaring variables
	var a int
	a = 1

	// one line
	var b int = 2

	// variable without specifying type, let the compiler figure it out
	c := 3

	fmt.Println("a:", a, ", b:", b, ", c:", c)

	var d float32 = 4
	// := could only declare float64 not float32
	// 5. specifies the variable would be a float instead of an int
	e := 5.
	fmt.Print("\n")
	fmt.Printf("%-3v %-10T %s", d, d, "\n")
	fmt.Printf("%-3v %-10T %s", e, e, "\n")
	fmt.Print("\n")

	// use variable blocks to declare multiple variables, keeps the code cleaner
	var (
		actorName    string = "Elisabeth Sladen"
		companion    string = "Sarah Jane Smith"
		doctorNumber int    = 3
		season       int    = 11
	)

	fmt.Printf("%s, %s, %v, %v %s", actorName, companion, doctorNumber, season, "\n")
	fmt.Print("\n")

	// variable could not be decalred twice in the same scope, but could be in different scopes
	// the variable in the inner-most scope takes precedence
	// this is called shadowing variable

	// changes the value of variable f from 5 to 6 (value 5 is assigned at package level)
	fmt.Println("The value of variable f at package scope is:", f)
	var f int = 6
	fmt.Println("The value of variable f at main scope is :", f)

	// using global variables
	fmt.Println("The value of global variable G is :", G)
	fmt.Print("\n")

	// converting variables with different types
	var h int = 8
	var i float32 = 8.5

	// typecast
	var j = float32(h)
	// loss of information when casting a float to an int
	// go does not go implicit typecast
	var k = int(i)

	fmt.Printf("%-3s %-10T %-5v %s", "h", h, h, "\n")
	fmt.Printf("%-3s %-10T %-5v %s", "i", i, i, "\n")
	fmt.Printf("%-3s %-10T %-5v %s", "j", j, j, "\n")
	fmt.Printf("%-3s %-10T %-5v %s", "k", k, k, "\n")

	// convert int to string (using string() will convert int number to the ascii number of char)
	var l int = 42
	var m string = string(l)
	fmt.Print("\nConvert int to string using string()\n")
	fmt.Printf("%-3s %-10T %-5v %s", "m", m, m, "\n")
	fmt.Print("Convert int to string using strconv.Iota()\n")
	m = strconv.Itoa(l)
	fmt.Printf("%-3s %-10T %-5v %s", "m", m, m, "\n")
}
