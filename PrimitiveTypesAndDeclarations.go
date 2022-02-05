package main

import "fmt"

func main() {
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

	/**
	COMPOSITE TYPES
	*/
	array1 := [...]uint{10, 20, 30}
	array2 := [12]int{41, 5: 60, 61, 10: 100}
	matrix := [2][3]int{}
	fmt.Printf("arrays %v, %v (length is part of the type of the array: %T) and matrix length %vx%v\n", array1, array2, array2, len(matrix), len(matrix[0]))

	slice1 := []uint{10, 20, 30}
	slice2 := make([]uint, 2, 3) // type, length, capacity
	slice1 = append(slice1, 40, 50)
	slice2 = append(slice1, slice2...)
	fmt.Printf("slice %v (Type %T, length %d, capacity %d) and slice2 %v \n", slice1, slice1, len(slice1), cap(slice1), slice2)
	slicedSlice := slice1[2:4]
	slicedSlice[0] = 555
	fmt.Printf("slice sliced %v (2,3) from %v\n", slicedSlice, slice1)
	toCopy := make([]uint, 3)
	numberCopied := copy(toCopy, slice1)
	toCopy[1] = 987
	fmt.Printf("slice copied %d elements. Original %v - New %v\n", numberCopied, slice1, toCopy)

	text := "Look, I'm typing Go"
	fmt.Printf("string is like an array, so substring '%s' from '%s' was easy\n", text[10:16], text)
	var rune1 rune = 'R'
	var byte1 byte = 178
	fmt.Printf("rune -> Str '%s' or byte -> Str '%s' or Str -> byte array %v\n", string(rune1), string(byte1), []byte(text))

}
