package main

import (
	"fmt"
	"os"
	"strconv"
)

// Requirements
// 1.Take in values for: width,length,no of parties
// 2.At least as many lots as parties
// 3.but no more than 10%
// 4.As square as possible

// Sollutions
//
//	easiest    not worth   hardest-best
//
// 1 - os.Args  || ReadLine || Config file
// 2 - Create a di
// 3 - at beggining calculate the 10%
// 4 -
// square is a shape with all sides equal - area = a^2 - perimiter = 4a - the ratio between 2 sides = 1
func main() {
	length, width, noParties, maxNoLots := getArguments()
	if length == 0 || width == 0 || noParties == 0 {
		return
	}
	fmt.Println(length, width, noParties, maxNoLots)
}

func getArguments() (uint8, uint8, uint8, uint8) {
	var length uint8
	var width uint8
	var noParties uint8
	var maxNoLots uint8

	arguments := os.Args[1:]

	if len(arguments) < 3 || arguments[0] == "help" {
		fmt.Println("Please run this binnary with the following argument order(Only numbers):")
		fmt.Println("length width numberOfParties")
		return 0, 0, 0, 0
	}
	// Check if only numeric
	res, err := strconv.Atoi(arguments[0])
	if err != nil {
		fmt.Println("Please input valid numbers")
		return 0, 0, 0, 0
	}
	length = uint8(res)

	res, err = strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Please input valid numbers")
		return 0, 0, 0, 0
	}
	width = uint8(res)
	res, err = strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Println("Please input valid numbers")
		return 0, 0, 0, 0
	}
	noParties = uint8(res)

	maxNoLots = noParties / 10
	return length, width, noParties, maxNoLots
}
