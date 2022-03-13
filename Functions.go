package main

import (
	"fmt"
	"sort"
)

func testFuctions() {
	fmt.Printf("I can now divide ('%d' / '%d') = '%d'\n", 781844, 64, divide(781844, 64))
	fmt.Printf("Variadic Func allow variable number of parameters: '%d', '%d'\n", variadicFunc(4, 1, 5, 87, 4), variadicFunc(1, 2, 3, 4, 5, 6, 7, 8, 9))

	a, b := 475, 29
	divided, modulo := multipleReturns(a, b)
	fmt.Printf("Multiple returns: ('%d' / '%d') = '%d' (mod %d)\n", a, b, divided, modulo)
	a, b = 1348, 61
	divided, modulo = multipleNamedReturns(a, b)
	fmt.Printf("Multiple named returns: ('%d' / '%d') = '%d' (mod %d)\n", a, b, divided, modulo)

	var functions = []func(int, int) int{divide, add}
	for idx, function := range functions {
		fmt.Printf("Functional programming result %d: ('%v')\n", idx, function(77817, 78))
	}

	for i := 88; i < 88*5; i += 88 {
		func(j int) {
			fmt.Printf("Inside anonymous function: %d\n", j)
		}(i)
	}

	sliceUnsorted := []string{"Annie", "Abed", "Jeff", "Pierce", "Dean", "Britta", "Shirley", "Troy", "Starburns"}
	fmt.Println("Unsorted list: ", sliceUnsorted)
	sort.Slice(sliceUnsorted, func(i, j int) bool {
		return len(sliceUnsorted[i]) > len(sliceUnsorted[j])
	})
	fmt.Println("Sorted list by passing a function: ", sliceUnsorted)

	deferFunc()
}

func divide(a int, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func add(a int, b int) int {
	return a + b
}

func multipleReturns(a int, b int) (int, int) {
	if b == 0 {
		return 0, 0
	}
	return a / b, a % b
}

func multipleNamedReturns(a int, b int) (divided int, modulo int) {
	if b == 0 {
		return divided, modulo
	}
	divided = a / b
	modulo = a % b
	return divided, modulo
}

func variadicFunc(numbers ...int) int {
	sum := 0
	for i, v := range numbers {
		sum += i * v
	}
	return sum
}

func deferFunc() {
	fmt.Println("Demo defer functionality that will always run after a function ends or errors occur")
	defer fmt.Println("DEFER STATEMENT")
	fmt.Println("Last line deferFunc()")
}
