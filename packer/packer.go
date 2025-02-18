package packer

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(input string) (string, error) {
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

		if count == 1 {
			resultStr.WriteRune(prevChar)
		} else if count > 1 { // исключаем символ с нулём
			resultStr.WriteString(strings.Repeat(string(prevChar), count))
		}

		if unicode.IsDigit(char) { // завершаем обработку пары "символ-цифра"
			prevChar = 0
		} else {
			prevChar = char
		}
	}

	if prevChar != 0 {
		resultStr.WriteRune(prevChar) // чтобы не забыть крайний символ
	}

	return resultStr.String(), nil
}

func Pack(input string) (string, error) {
	// fmt.Printf("%s\n", a1)

	var count int
	resultStr := strings.Builder{}
	var prevChar rune // предыдущий символ исходной строки

	for _, char := range input { // разбираем входящую строку на символы
		// fmt.Printf("%c\n", i)

		if prevChar == 0 { // init
			if unicode.IsDigit(char) {
				return "", errors.New("ожидается символ")
			}
			prevChar = char
			count = 1
			continue
		}

		if unicode.IsDigit(char) { // проверка цифры
			return "", errors.New("ожидается символ")
		}

		if char == prevChar {
			count++
		} else {
			resultStr.WriteRune(prevChar)

			if count > 1 {
				resultStr.WriteString(strconv.Itoa(count))
			}
			prevChar = char
			count = 1
		}
	}

	if prevChar != 0 {
		resultStr.WriteRune(prevChar) // чтобы не забыть крайний символ
	}
	if count > 0 {
		resultStr.WriteString(strconv.Itoa(count))
	}

	return resultStr.String(), nil
}
