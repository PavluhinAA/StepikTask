package main

import (
	"fmt"
)

func exit() {
	fmt.Println("\nВведите любой символ чтобы выйти")
	var exit string
	fmt.Scan(&exit)
}
func division(x, y float64) float64 {
	res := x / y
	return res
}
func multiplication(x, y float64) float64 {
	res := x * y
	return res
}
func subtraction(x, y float64) float64 {
	res := x - y
	return res
}
func addition(x, y float64) float64 {
	res := x + y
	return res
}
func find_perc(x, y float64) float64 {
	res := y / x * 100
	return res
}
func calculator(x, y float64, input string) float64 {
	var res float64
	switch input {
	case `*`:
		res = multiplication(x, y)
	case "/":
		if y == 0 {
			fmt.Println("Деление на ноль невозможно")
			panic("")
		} else {
			res = division(x, y)
		}
	case "+":
		res = addition(x, y)
	case "-":
		res = subtraction(x, y)
	case "%":
		res = find_perc(x, y)
	default:
		fmt.Println("Неизвестный знак действия")
	}
	return res
}
func main() {
	defer exit()
	var (
		input     string
		x, y, res float64
	)
	fmt.Println("Введите запрос")
	_, err := fmt.Scan(&x, &input, &y)
	if err != nil {
		fmt.Println("Проверьте правильность ввода")
		panic("")
	}
	res = calculator(x, y, input)

	if float64(int(res)*3) == res*3 {
		fmt.Printf("Ваш результат %.0f", res)
	} else if float64(int(res*10)) == res*10 {
		fmt.Printf("Ваш результат %.1f", res)
	} else {
		fmt.Printf("Ваш результат %.2f", res)
	}
}

// Как тебе мой велосипед?
