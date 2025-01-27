package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	var a1 string

	for {
		_, err := fmt.Fscanln(os.Stdin, &a1)
		if err != nil {
			panic("input err: " + err.Error())
		}
		// fmt.Println(a1)
		fmt.Printf("%q\n", unpack(a1))
	}

	a := "f3я1х2fa4by0c2\n3d5t"
	s := unpack(a)

	fmt.Printf("%q\n", s)
}

func unpack(a1 string) string {

	var s string // результирующая строка
	var p string // предыдущий символ исходной строки

	for _, i := range a1 {

		if p == "" { // проверка буквы из каждой пары "символ-цифра"
			if unicode.IsDigit(i) {
				panic("ожидается буква")
			}
			p = string(i)
			continue
		}

		count := 1 // если цифра отсутствует

		if unicode.IsDigit(i) { // проверка цифры из каждой пары
			digit, err := strconv.Atoi(string(i))
			if err != nil {
				panic("parse error")
			}
			fmt.Printf("%d is Number\n", digit)
			count = digit
		}

		if count > 0 { // исключаем символ с нулём
			s += strings.Repeat(p, count)
		}

		if unicode.IsDigit(i) { // завершаем обработку пары "символ-цифра"
			p = ""
		} else {
			p = string(i)
		}
	}
	s += p // чтобы не забыть крайний символ

	return s
}
