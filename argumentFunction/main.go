package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evenSum(a, b int) {
	if a > b {
		a, b = b, a
	}

	sum := 0
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	fmt.Printf("Сумма всех чётных чисел в диапозоне от %d до %d равна: %d \n", a, b, sum)
}

func main() {
	fmt.Println("Введите два числа, а я вычислю сумму чётных чисел в этом диапозоне!")

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

	evenSum(numOne, numTwo)
}
