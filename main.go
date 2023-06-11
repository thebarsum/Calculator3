package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rim = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
	"VII": 7, "VIII": 8, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90, "C": 100}
var convRim = [14]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 40, 50, 90, 100}
var a, b *int
var matOps = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var data []string

func main() {
	fmt.Println("Input:")
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(text, " ", "")
		calk(strings.ToUpper(strings.TrimSpace(s)))
	}
}
func calk(s string) {
	var matOp string
	var stringsFound int
	numbers := make([]int, 0)
	rims := make([]string, 0)
	rimToInt := make([]int, 0)
	for idx := range matOps {
		for _, val := range s {
			if idx == string(val) {
				matOp += idx
				data = strings.Split(s, matOp)
			}
		}
	}
	switch {
	case len(matOp) > 1:
		panic("Больше 1 оператора")
	case len(matOp) < 1:
		panic("Нет оператора")
	}
	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			rims = append(rims, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 1:
		panic("Разные системы счисления")
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 &&
			numbers[1] > 0 && numbers[1] < 11
		if val, ok := matOps[matOp]; ok && errCheck == true {
			a, b = &numbers[0], &numbers[1]
			fmt.Println("Output:")
			fmt.Println(val())
		} else {
			panic("Интервал чисел от 0 до 10 включительно")
		}
	case 2:
		for _, elem := range rims {
			if val, ok := rim[elem]; ok && val > 0 && val < 11 {
				rimToInt = append(rimToInt, val)
			} else {
				panic("Интервал чисел от 1 до 10 включительно")
			}
		}
		if val, ok := matOps[matOp]; ok {
			a, b = &rimToInt[0], &rimToInt[1]
			intToRoman(val())
		}
	}
}
func intToRoman(rimResult int) {
	var ReConvRim []int
	var rimNum string
	if rimResult == 0 {
		panic("В римской системе нет 0")
	} else if rimResult < 0 {
		panic("В римской системе нет отрицательныйх чисел")
	}
	for i := len(convRim) - 1; i >= 0; i-- {
		ReConvRim = append(ReConvRim, convRim[i])
	}
	for rimResult > 0 {
		for _, elem := range ReConvRim {
			for i := elem; i <= rimResult; {
				for index, value := range rim {
					if value == elem {
						rimNum += index
						rimResult -= elem
					}
				}
			}
		}
	}
	fmt.Println("Output:\n" + rimNum)
}
