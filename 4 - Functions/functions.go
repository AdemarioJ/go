package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculateMath(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	sum := sum(5, 3)
	fmt.Println(sum)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Text function")
	fmt.Println(result)

	resultSum, resultSub := calculateMath(5, 3)
	fmt.Println(resultSum, resultSub)

	sum, _ = calculateMath(10, 5)
	fmt.Println(sum)
}
