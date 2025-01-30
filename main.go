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

	/**
	Вариант 1. Ручной ввод строки в консоли
	*/
	for {
		fmt.Print("Введите строку: ")
		_, err := fmt.Fscanln(os.Stdin, &a1)
		if err != nil {
			if err.Error() == "unexpected newline" {
				continue
			}
			panic("input err: " + err.Error())
		}

		fmt.Printf("%q\n", unpack(a1))
	}

	/**
	Вариант 2. Строка задана в переменной
	*/
	// a1 = "f3я1х2fa4by0c2\n3d5t"
	// fmt.Printf("%q ==> \n%q\n", a1, unpack(a1))
}

func unpack(a1 string) string {

	var s string // результирующая строка
	var p string // предыдущий символ исходной строки

	for _, i := range a1 {

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
