// Package represent logical grouping of multiple go files.
// A folder can only have one package in it and a package can only span over one directory
// If a package contains runnable entry point then the package has to be named main, otherwise a more complex naming schema is followed where last part of package name has to match the folder name
package main

// import statement allows you to import other packages to be used in your application
import (
	"fmt"
	"math/rand"
	"time"
)

// main function in main package is the entry point for the applicartion
// A package can only have one main function
func main() {

	// Variable declaration syntax
	// var s1 string
	var s1 = "This is first string"
	fmt.Println(s1)
	s1 = "First String"
	fmt.Println(s1)

	// s2 = "Second String"
	s2 := "Second String"
	fmt.Println(s2)

	// Basic Types

	// string default value ""
	var tstS1 string
	fmt.Println(tstS1 == "")

	// bool default value false
	var tstBool1 bool
	fmt.Println(tstBool1)

	// numeric types
	// integer types int8, uint8 (byte), int16, uint16, int32 (rune), uint32, int64, uint64, int, uint, uintptr
	// decimal types float32, float64
	// complex types complex64, complex128

	var tstInt1 int
	fmt.Println(tstInt1)

	// Flow Control

	// For
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	check := true

	for check {
		fmt.Println("in go for acts like a while in other languages")
		check = false
	}

	var tstArr [1]string
	strVal := "test String in array"
	tstArr[0] = strVal

	// var v is by value so its not the same variable as tstArr[i]
	for i, v := range tstArr {
		fmt.Printf("%d: %s\t%p\t%p\n", i, v, &v, &tstArr[i])
	}

	for i := range tstArr {
		fmt.Printf("%d: %s\t%p\n", i, tstArr[i], &tstArr[i])
	}

	// if

	check = true

	if check {
		fmt.Println("default `if` syntax")
	}

	if check2 := true; check2 {
		fmt.Println("short syntax for if")
	}

	// switch statement
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday.")
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Today is a weekday.")
	}

	switch num := rand.Intn(100); {
	case num < 50:
		fmt.Println("random number less than 50")
	case num > 50 && num < 80:
		fmt.Println("random number greater than 50 but less than 80")
		// fallthrough
	default:
		fmt.Println("random number greater than 80")
	}
}