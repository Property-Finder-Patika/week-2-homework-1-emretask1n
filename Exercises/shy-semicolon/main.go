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
// EXERCISE: Shy Semicolons
//
//  1. Try to type your statements by separating them using
//     semicolons
//
//  2. Observe how Go fixes them for you
//
// ---------------------------------------------------------

func main() {

	//  Gofmt tool will formats our code automatically

	/*
		fmt.Println("inanc"); fmt.Println("lina"); fmt.Println("ebru");
	*/

	// this code turns into that

	fmt.Println("inanc")
	fmt.Println("lina")
	fmt.Println("ebru")
}
