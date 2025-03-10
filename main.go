package main

import "fmt"

func main() {

	fact := factorial(5)
	fmt.Println(fact)

	sum := variadicFn(1, 15, 20, 50, 10)

	numbers := []int{15, 20, 50, 10}
	anotherSum := variadicFn(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func factorial(number int) int {
	if number == 1 {
		return 1
	}
	return number * factorial(number-1)
	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result
}

func variadicFn(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
