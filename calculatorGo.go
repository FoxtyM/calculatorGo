package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numRim map[string]int
var res int
var arab []int
var convInttoRoman = [9]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
var a, b int
var matchedArab, matchedRoman bool
var FirstNumber, SecondNumber int

func main() {
	numRim = map[string]int{"C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
	//Чтение введенных данных из консоли
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()
	result := strings.Split(str, " ")

	matchedArab, _ = regexp.MatchString(`^\d{1,2}\s[+-/*]\s\d{1,2}$`, str)
	matchedRoman, _ = regexp.MatchString(`^[IVX]{1,4}\s[+-/*]\s[IVX]{1,4}$`, str)

	if matchedArab {
		a, err := strconv.Atoi(result[0])
		if err != nil || a <= 0 || a > 10 {
			panic("Ошибка, калькулятор принимает на вход числа от 1 до 10 включительно")
		}
		b, err := strconv.Atoi(result[2])
		if err != nil || b <= 0 || b > 10 {
			panic("Ошибка, калькулятор принимает на вход числа от 1 до 10 включительно")
		}
		fmt.Println(calculate(result, a, b))
	} else if matchedRoman {
		romantoArab(result)
		calculate(result, a, b)
		fmt.Println(arabtoRoman(res))
	} else {
		panic("Ошибка, формат математической операции не удовлетворяет заданию — два одинаковых операнда (либо только римских, либо только арабских) и один оператор (+, -, /, *).")
	}
}

func romantoArab(result []string) []int {

	for _, value := range result {
		if numRim[value] != 0 {
			arab = append(arab, numRim[value])
		} else if len(value) > 1 {
			str1 := strings.Split(value, "") //Получаем V I I I
			ints := make([]int, len(str1))
			for index, value := range str1 {
				ints[index] = numRim[value]
			}
			// 5 1 1 1 []int
			sumNum := 0
			for _, value := range ints {
				sumNum += value
			}
			arab = append(arab, sumNum)
		}
	}
	return arab
}

func calculate(result []string, a, b int) int {

	if matchedRoman {
		FirstNumber = arab[0]
		SecondNumber = arab[1]

		//Проверка простого или сложносоставного римского числа
		CheckFirstNumber := arabtoRoman(arab[0])
		CheckSecondNumber := arabtoRoman(arab[1])
		F, _ := regexp.MatchString(CheckFirstNumber, result[0])
		S, _ := regexp.MatchString(CheckSecondNumber, result[2])

		if !F || !S {
			panic("Ошибка, неверное написание одного из двух римских операндов")
		}

		if arab[0] <= 0 || arab[0] > 10 || arab[1] <= 0 || arab[1] > 10 {
			panic("Ошибка, калькулятор принимает на вход числа от 1 до 10 включительно")
		}
	} else if matchedArab {
		FirstNumber = a
		SecondNumber = b
	}

	switch result[1] {
	case "+":
		res = FirstNumber + SecondNumber
	case "-":
		res = FirstNumber - SecondNumber
	case "*":
		res = FirstNumber * SecondNumber
	case "/":
		res = FirstNumber / SecondNumber
	}

	return res
}

func arabtoRoman(res int) string {
	var resRoman []string
	if res <= 0 {
		panic("Ошибка, в римской системе нет 0 и отрицательных чисел")
	} else {
		for res > 0 {
			for _, v := range convInttoRoman {
				for a := v; a <= res; {
					for key, value := range numRim {
						if value == v {
							resRoman = append(resRoman, key)
							res = res - a
						}
					}
				}
			}
		}
		return strings.Join(resRoman, "")
	}
}
