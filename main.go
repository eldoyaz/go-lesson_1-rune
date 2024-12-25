package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	const r1 = '€'
	fmt.Println("(int32) r1:", r1)
	fmt.Printf("(HEX) r1: %x\n", r1)
	fmt.Printf("(as a String) r1: %s\n", r1)
	fmt.Printf("(as a character) r1: %c\n", r1)

	/*aString := []byte("Miхаlis")
	// aString := "Miхаlis"
	for x, y := range aString {
		fmt.Println(x, y, unicode.IsSymbol(rune(y)))
		fmt.Printf("Elem[x]: %c, Char y: %c\n", aString[x], y)
	}
	fmt.Printf("%s\n", aString)

	return*/

	a1 := `fяfa44by0c2\n3d5e`
	a2 := "fяхfa44by0c2\n3d5e"

	fmt.Printf("%q %d\n", a1, len(a1))
	fmt.Printf("%q %d\n", a2, len(a2))

	var b string // результирующая строка
	var c string // текущий символ исходной строки

	for i, a := range a2 {
		fmt.Printf("%d => (int)%v; char => %q\n", i, a, a)

		if !unicode.IsDigit(a) {
			fmt.Printf("%q not Number\n", a)
			c = string(a)
			// b += string(a)
			// n = true
		} else {
			n, err := strconv.Atoi(string(a))
			if err != nil {
				panic("parse error")
			}
			fmt.Printf("%d a Number\n", n)
			b += strings.Repeat(c, n)
		}
	}
	fmt.Printf("%q", b)

	return
}

func unpack(a string) string {

	r, err := regexp.Compile("[[:alpha:][:space:]]{1}[[:digit:]]?")
	if err != nil {
		log.Fatal("regexp compile error")
	}
	s := r.FindAllStringSubmatch(a, 100)
	fmt.Println(s)

	var c string
	for _, s2 := range s {
		for _, s3 := range s2 {
			fmt.Printf("%#v\n", s3)
			c += decode(s3)
		}
	}

	return c
}

func unpack2(a string) string {

	// var b []rune
	// var i int

	for _, c := range a {
		fmt.Printf("%c --> %d\n", c, c)
		fmt.Printf("%c is digit == %t\n", c, unicode.IsDigit(c))
	}

	return a
}

func decode(s string) string {
	var c []byte
	var d int = 1

	if len(s) > 1 && unicode.IsDigit(rune(s[1])) {
		d, _ = strconv.Atoi(string(s[1]))
	}

	for i := 0; i < d; i++ {
		c = append(c, s[0])
	}

	return string(c)
}

func hello() {

	s := "Hello, Мир!" // 11 -> 14

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i]) // Hello, ABCDEF!
	}

	fmt.Println()

	for _, r := range s {
		fmt.Printf("%c ", r) // Hello, Мир!
	}

	fmt.Println()

	for i, _ := range s { //
		fmt.Printf("%c ", s[i]) // Hello, ACE!
	}

	fmt.Println()
}
