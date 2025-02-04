package packer

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(a1 string) string {

	var s string // результирующая строка
	var p string // предыдущий символ исходной строки

	for _, i := range a1 {
		fmt.Printf("%q\n", i)

		if p == "" { // проверка буквы из каждой пары "символ-цифра"
			if unicode.IsDigit(i) {
				println("некорректная строка (ожидается символ)")
				return ""
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
			// fmt.Printf("%d is Number\n", digit)
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

func Pack(a1 string) string {

	var count int
	var s string // результирующая строка
	var p string // предыдущий символ исходной строки

	for _, i := range a1 {
		// fmt.Printf("%q\n", i)

		if unicode.IsDigit(i) {
			println("некорректная строка (ожидается символ)")
			return ""
		}

		if p == "" { // init
			p = string(i)
			count = 1
			continue
		}

		if string(i) == p {
			count++
		} else {
			s += p
			if count > 1 {
				s += strconv.Itoa(count)
			}

			p = string(i)
			count = 1
		}
	}

	s += p // чтобы не забыть крайний символ
	if count > 1 {
		s += strconv.Itoa(count)
	}

	return s
}
