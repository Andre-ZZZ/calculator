package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intToRoman(i int) string {
	arabicRoman := []struct {
		a int
		r string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	res := ""
	for _, v := range arabicRoman {
		for i >= v.a {
			res += v.r
			i -= v.a
		}
	}
	return res
}

func doOp(a, b int, op uint8) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	default:
		panic("unknown op")
	}
}

var roman = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
var arabic = map[string]int{
	"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10}

func process(input string) (string, error) {
	op := strings.IndexAny(input, "+-/*")
	if op == -1 {
		return "", fmt.Errorf("строка не является арифметической операцией")
	}
	if op != strings.LastIndexAny(input, "+-/*") {
		return "", fmt.Errorf("в строке допустима только одна арифметическая операция")
	}
	l := strings.TrimSpace(input[:op])
	r := strings.TrimSpace(input[op+1:])

	a := roman[l] + arabic[l]
	if a == 0 {
		return "", fmt.Errorf("первый операнд не удовлетворяет заданию")
	}
	b := roman[r] + arabic[r]
	if b == 0 {
		return "", fmt.Errorf("второй операнд не удовлетворяет заданию")
	}

	res := doOp(a, b, input[op])
	if roman[l] > 0 && roman[r] > 0 { // два римских числа
		if res < 1 {
			return "", fmt.Errorf("в римской системе нет отрицательных чисел и нуля")
		}
		return intToRoman(res), nil
	} else if arabic[l] > 0 && arabic[r] > 0 { // два арабских числа
		return strconv.Itoa(res), nil
	} else {
		return "", fmt.Errorf("используются одновременно римские и арабские")
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		//fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		res, err := process(text)
		if err != nil {
			fmt.Print("Вывод ошибки, так как ")
			fmt.Println(err)
			return
		}
		fmt.Println(res)
	}
}
