package packer

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(a1 string) string {
	// fmt.Printf("%s\n", a1)

	var s string // результирующая строка
	var p string // предыдущий символ исходной строки

	for _, i := range a1 {
		// fmt.Printf("%c\n", i)

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

		if p == "" { // init
			p = string(i)
			count = 1
			continue
		}

		if unicode.IsDigit(i) {
			println("некорректная строка (ожидается символ)")
			return ""
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

func Unpack1(input string) (string, error) {
	// fmt.Printf("%s\n", a1)

	resultStr := strings.Builder{}
	var prevChar rune // предыдущий символ исходной строки

	for _, char := range input { // разбираем входящую строку на символы
		// fmt.Printf("%c\n", i)

		if prevChar == 0 { // проверка "ведущей буквы" из каждой пары "символ-цифра"
			if unicode.IsDigit(char) {
				return "", errors.New("ожидается символ")
			}
			prevChar = char
			continue
		}

		count := 1 // если цифра отсутствует

		if unicode.IsDigit(char) { // проверка цифры из каждой пары
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				return "", errors.New("digit parse error")
			}
			// fmt.Printf("%d is Number\n", digit)
			count = digit
		}

		if count > 0 { // исключаем символ с нулём
			resultStr.WriteString(strings.Repeat(string(prevChar), count))
		}

		if unicode.IsDigit(char) { // завершаем обработку пары "символ-цифра"
			prevChar = 0
		} else {
			prevChar = char
		}
	}
	resultStr.WriteRune(prevChar) // чтобы не забыть крайний символ

	return resultStr.String(), nil
}
