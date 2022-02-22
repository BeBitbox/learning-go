package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func blocksShadowsContrulStructures() {
	shadowing()
	ifLoop()
	forLoop()
	switchStuff()
}

func shadowing() {
	x := 10
	if x > 5 {
		x := 4
		fmt.Printf("shadow %v", x)
	}
	fmt.Printf(" but after the controlblock %v and shawdowing '%v' ", x, true)
	true := "false"
	fmt.Printf("becomes '%v' because true is just an identifier in the universal block\n", true)
}

func ifLoop() {
	// always 1 apparently. Variable n only exists in the block of the if/else
	if n := rand.Intn(10); n < 3 {
		fmt.Printf("Random n(%v) is smaller than three\n", n)
	} else if n < 7 {
		fmt.Printf("Random n(%v) is smaller than seven\n", n)
	} else {
		fmt.Printf("Random n(%v) is bigger than six\n", n)
	}
}

func forLoop() {
	text := "complete for train "
	for i := 0; i < 8; i++ {
		text += "-> " + strconv.Itoa(i)
	}
	fmt.Println(text)

	text = "condition only for Fibonnaci "
	a := 1
	b := 2
	for b < 1000 {
		text += "-" + strconv.Itoa(a) + "-" + strconv.Itoa(b)
		a = a + b
		b = a + b
	}
	fmt.Println(text)

	text = "Infinte for "
	for {
		n := rand.Intn(10)
		text += " -> " + strconv.Itoa(n)
		if n == 2 {
			break
		}
	}
	fmt.Println(text)

	text = "Ranged for "
	slice := []int{22, 34, 4, 947, 4, 5}
	for i, v := range slice {
		text += "(" + strconv.Itoa(i) + "-" + strconv.Itoa(v) + ")"
	}
	fmt.Println(text)
}

func switchStuff() {
	words := []string{"Gimme", "Abba", "QueenS", "Dancing", "Waterloo", "Money"}
	for i, word := range words {
		switch n := len(word); n {
		case 3:
			fmt.Printf("3 letters on %d: %s\n", i, word)
		case 4, 5:
			fmt.Printf("4 or 5 letters on %d: %s\n", i, word)
		case 7:
			fmt.Printf("7 letters on %d: %s\n", i, word)
		default:
			fmt.Printf("Other number of letters letters on %d: %s\n", i, word)

		}
	}

statement:
	for _, word := range words {
		switch {
		case len(word) < 1:
			fmt.Printf("Small word found")
		case len(word) > 6:
			fmt.Printf("5+ letters word: %s\n", word)
			break statement
		}
	}
}
