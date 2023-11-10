package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inverting(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println("Введите два числа, а я поменяю их местами!")

	reader := bufio.NewReader(os.Stdin)

	var numOne int
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var err error
		numOne, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите корректное целое число.")
			continue
		}
		break
	}

	var numTwo int
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var err error
		numTwo, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите корректное целое число.")
			continue
		}
		break
	}

	fmt.Printf("Ваши числа: %d, %d\n", numOne, numTwo)

	inverting(&numOne, &numTwo)

	fmt.Printf("Перевернутые числа: %d, %d\n", numOne, numTwo)
}
