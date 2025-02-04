package main

import (
	"errors"
	"fmt"
	"os"

	"lesson_2_rune/packer"
)

var newLineErr = errors.New("expected newline")
var noNewLineErr = errors.New("unexpected newline")

func main() {

	var flag1 string
	if len(os.Args) > 1 {
		flag1 = os.Args[2]
	}
	fmt.Println("flag", flag1)
	// return

	var fn1 func(string) string
	switch flag1 {
	case "--unpack", "--daemon", "--input":
		fn1 = packer.Unpack
		println("режим Распаковки")
	case "--pack":
		fn1 = packer.Pack
		println("режим Упаковки")
	default:
		println("некорректный аргумент")
		println("ожидается: --pack, --unpack, --daemon, --input q1w2e3")
		return
	}
	_ = fn1("")

	var a1 string

	/** _________________________________
	Вариант 1. Строка задана в переменной
	*/
	// a1 = "f3я\\2х3fa4by0c2\n3d5t"
	// // a1 = `f3я\2х3\\nfa4by0c2\\n3d\5t` // @todo
	// fmt.Printf("%q ==> \n%q\n", a1, packer.Unpack(a1))

	/** _________________________________
	Вариант 2. Строка задана в переменной
	*/
	// a1 = "a_bbb\n\n\t   "
	// fmt.Printf("%q ==> \n%q\n", a1, packer.Pack(a1))

	/** ___________________________________
	Вариант 3. Ручной ввод строки в консоли
	*/
	for {
		fmt.Print("Введите строку: ")
		_, err := fmt.Fscanln(os.Stdin, &a1)
		if err != nil {
			if err.Error() == newLineErr.Error() {
				continue
			}
			if err.Error() == noNewLineErr.Error() {
				continue
			}
			panic("input err: " + err.Error())
		}

		fmt.Printf("%q\n", fn1(a1))
	}
}
