package main

import (
	"fmt"
	"unicode"
)

func main() {
	// "a4bc2d5e" => "aaaabccddddde"

	a := "a4bc2d5e"
	b := unpack(a)
	fmt.Println(b)
}

func unpack(a string) string {

	for i, c := range a {
		fmt.Printf("%c --> ", c)
		println(a[i])
		fmt.Printf("%c is digit == %t\n", c, unicode.IsDigit(c))
	}

	return a
}
