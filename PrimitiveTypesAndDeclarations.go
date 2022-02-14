package main

import "fmt"

func primitivesAndDeclarations() {
	var binary = 0b100
	var octal = 0100
	var decimal = 100
	var hexa = 0x100
	fmt.Printf("printing numbers with different base: 0100 becomes binary(%d), octal(%d), decimal(%d) hex(%d)\n", binary, octal, decimal, hexa)

	var runeLiteral = 'a'
	var runeLiteralHex = '\uA970'
	var runeLiteralUnicode = '\U0000F970'
	fmt.Printf("printing runeLiterals %c %c %c\n", runeLiteral, runeLiteralHex, runeLiteralUnicode)
	var string1 = "Tractor"
	var string2 = `says
		vroem`
	fmt.Printf("printing String literals %s %s\n", string1, string2)

	var boolean1 bool
	var boolean2 = true || string1 == "Tr"+"actor"
	fmt.Printf("printing Booleans %t en  %t\n", boolean1, boolean2)

	var x64 uint64 = 144_454_547_874_747
	var a int64 = 21
	var b = int32(a / 5)
	var a2 float32 = 1.45e10
	complexNumber := complex(87878, 48484)
	fmt.Printf("printing %d / 5 = %d\n", a, b)
	fmt.Printf("printing long numbers with underscores %d\n", x64)
	fmt.Printf("printing float32 numbers with exponents %f\n", a2)
	fmt.Printf("printing %T numbers %v\n", complexNumber, complexNumber)

	var (
		x1 int
		x2 = 45
		s1 string
	)
	fmt.Printf("printing declartionList numbers %v-%v-%v\n", x1, x2, s1)

	const constantSum = 78 + 7882 + 4
	const constantString string = "ahha"
	const constantComplex = complex(717, 77)
	fmt.Printf("printing constant variable %v %v %v\n", constantSum, constantString, constantComplex)
}
