package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romeToArab map[string]int = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VII":  6,
	"VIII": 7,
	"IIX":  8,
	"IX":   9,
	"X":    10,
}

var arabToRome map[int]string = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}

var arabToRomeDec map[int]string = map[int]string{
	1: "X",
	2: "XX",
	3: "XXX",
	4: "XL",
	5: "L",
	6: "LX",
	7: "LXX",
	8: "LXXX",
	9: "XC",
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите выражение: ")

		// Разделение считываемой строки на 3 группы символов,
		// в качестве разделителя используется пробел
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")

		// Валидация ввода
		val1, isRomeVal1 := toInt(parts[0])
		// То же самое и для второго входного значения
		val2, isRomeVal2 := toInt(parts[2])
		// Чтение оператора
		operator := parts[1]

		// Одним из условий является то, что на входе числа в одинаковом формате
		if isRomeVal1 != isRomeVal2 {
			panic("Числа в разном офрмате!")
		}
		// Одним из условий является то, что на входе числа не более 10 включительно
		if val1 > 10 || val2 > 10 {
			panic("Введенные значения больше 10!")
		}

		switch operator {
		case "+":
			if isRomeVal1 {
				fmt.Println(toRome(val1 + val2))
			} else {
				fmt.Println(val1 + val2)
			}
		case "-":
			// Одно из условий, что результат выражения в римском формате не может быть отрацательным
			if isRomeVal1 {
				if val1-val2 < 1 {
					panic("Выражение в римском формате и результат меньше единицы!")
				}
			}

			if isRomeVal1 {
				fmt.Println(toRome(val1 - val2))
			} else {
				fmt.Println(val1 - val2)
			}
		case "*":
			if isRomeVal1 {
				fmt.Println(toRome(val1 * val2))
			} else {
				fmt.Println(val1 * val2)
			}
		case "/":
			if isRomeVal1 {
				fmt.Println(toRome(val1 / val2))
			} else {
				fmt.Println(val1 / val2)
			}
		default:
			panic("Не подходящая арифметитческая операция!")
		}
	}
}

// Если на вход подается арабское число, то Atoi его конвертирует
// Если вход есть римское число, то Atoi выдаст ошибку
// Затем из-за ошибки проверяем хэштаблицу на предмет
// наличия в ней подобной комбинации символов
// Если в хэштаблице нет такой комбинации, то вызываем панику
func toInt(s string) (val int, isRome bool) {
	val, err := strconv.Atoi(s)
	if err != nil {
		var ok bool
		val, ok = romeToArab[s]
		isRome = true
		if ok != true {
			panic("Неверный формат входных данных")
		}
	}
	return
}

func toRome(x int) (str string) {
	if x == 100 {
		return "C"
	} else if x < 100 && x > 10 {
		return arabToRomeDec[x/10] + arabToRome[x%10]
	} else {
		return arabToRome[x]
	}
}
