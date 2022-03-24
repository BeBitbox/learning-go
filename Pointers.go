package main

import (
	"fmt"
)

func pointers() {
	x := 6894
	pointer := &x
	fmt.Printf("Show pointer of %v: %v, and dereference it to %v\n", x, pointer, *pointer)

	var nilPointer *int
	fmt.Printf("Nil pointer %v: %t\n", nilPointer, nilPointer)

	//  var strPointer *string = &"aaarrar"
	var strPointer *string = depointer("aaarrar")
	fmt.Printf("You can't create pointer from constant, use function %v: %v\n", strPointer, *strPointer)
}

func depointer(s string) *string {
	return &s
}
