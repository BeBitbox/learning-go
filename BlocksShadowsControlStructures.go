package main

import "fmt"

func blocksShadowsContrulStructures() {
	shadowing()
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
