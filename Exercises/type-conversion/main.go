// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert and Fix
//
//  Fix the codes by using the conversion expression.
//
// EXPECTED OUTPUT-1
//  15.5

// EXPECTED OUTPUT-2
//  10.5

// EXPECTED OUTPUT-3
//  5.5

// EXPECTED OUTPUT-4
//  9.5

// EXPECTED OUTPUT-5
//  1127
// ---------------------------------------------------------

func main() {
	// a, b := 10, 5.5
	// fmt.Println(a + b)
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)

	// a, b := 10, 5.5
	// a = b
	// fmt.Println(a + b)
	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)

	// fmt.Println(int(5.5))
	fmt.Println(5.5)

	// age := 2
	// fmt.Println(int(7.5) + int(age))
	age := 2
	fmt.Println(7.5 + float64(age))

	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	// fmt.Println(int8(max) + min)
	fmt.Println(max + int16(min))
}
