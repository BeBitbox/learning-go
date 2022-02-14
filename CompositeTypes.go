package main

import "fmt"

func compositeTypes() {
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

	totalWins := map[string]int{
		"Jos":   5,
		"Jef":   0,
		"Venus": 1,
	}
	totalWins["Jan"] = 1
	totalWins["Jef"] = 3
	delete(totalWins, "Venus")
	fmt.Printf("Maps %v with length %d\n", totalWins, len(totalWins))

	type person struct {
		name string
		age  uint8
	}
	betty := person{
		name: "Betty",
		age:  45,
	}
	janinneke := person{"Janinneke", 30}
	janinneke.age = 130
	cat := struct {
		legs  uint8
		sound string
	}{
		4,
		"Miauw",
	}
	fmt.Printf("Persons Betty: %v, Janinneke: %v and the anonymous struct %v\n", betty, janinneke, cat)
}
