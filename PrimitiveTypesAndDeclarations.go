package main

import "fmt"

func main() {
	var binary = 0b100
	var octal = 0100
	var decimal = 100
	var hexa = 0x100
	fmt.Printf("printing numbers with different base: 0100 becomes binary(%d), octal(%d), decimal(%d) hex(%d)\n", binary, octal, decimal, hexa)
	fmt.Printf("printing long numbers with underscores %d\n", 123_456_789)
	fmt.Printf("printing float numbers with exponents %f\n", 1.45e10)
	var runeLiteral = 'a'
	var runeLiteralHex = '\uA970'
	var runeLiteralUnicode = '\U0000F970'
	fmt.Printf("printing runeLiterals %c %c %c\n", runeLiteral, runeLiteralHex, runeLiteralUnicode)
	var string1 = "Tractor"
	var string2 = `says
		vroem`
	fmt.Printf("printing String literals %s %s\n", string1, string2)

}
