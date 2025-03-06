package functionsasvalues

import "fmt"

type transformFn func(int) int // moze se koristiti unutar transformNumbers ali meni je dole citljivije

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := tranformNumbers(&numbers, double)
	tripled := tranformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	tranformedNumbers := tranformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := tranformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(tranformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func tranformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}
