package main //   go run main.go

import (
	"fmt"
	"strings"
)

func main() {
	var maps = map[string]int{ //задание 1
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	fmt.Println("Добавть нового человека")
	var key_new string
	var value_new int
	fmt.Scan(&key_new, &value_new)
	maps[key_new] = value_new
	fmt.Println(maps)

	map_return(maps) // задание 2

	fmt.Println("Удалите человека из жизни: ") // Задание 3
	var del string
	fmt.Scan(&del)
	delete(maps, del)
	fmt.Println("Вот список без трупа: ", maps)

	fmt.Println("Введите строку: ") //  задание 4
	var up_string string
	fmt.Scan(&up_string)
	fmt.Println(strings.ToUpper(up_string))

	fmt.Println("Введите несколько чисел: ") //задание 5
	var arr []int
	sum := 0
	i := 0
	for {
		fmt.Scan(&i)
		arr = append(arr, i)
		sum += i
		if i == 0 {
			break
		}
		i++
	}
	fmt.Println("Сумма: ", sum)

	fmt.Println("Задание 6") //задание 6
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}
func map_return(average_value map[string]int) { // функция задания 2
	var a int
	for key, _ := range average_value {
		a += average_value[key]
	}
	a /= len(average_value)
	fmt.Println("Среднее значение возростов: ", a)
}

//   go run main.go
