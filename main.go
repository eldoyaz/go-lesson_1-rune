package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	// "a4bc2d5e" => "aaaabccddddde"

	a := "ffa44by0c2\n2d5e"

	b := unpack(a)
	fmt.Println(b)
	fmt.Printf("%#v\n", b)

	// c := unpack2(a)
	// fmt.Println(c)
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
