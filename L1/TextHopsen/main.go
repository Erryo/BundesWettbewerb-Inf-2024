package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const PATH_TO_FILE string = "../43.1/J2_Texthopsen/hopsen1.txt"

// Indicates who should win in case of EOF error
// false - player 1
// true - player 2
var whoseJump bool

func main() {
	//	for i := range 6 {
	//		var PATH_TO_INPUT string = fmt.Sprintf("../43.1/J1_QuadratischPraktischGrБn/garten%v.txt", i)
	//	}

	germanAlphabetMap := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
		"ä": 27, "ö": 28, "ü": 29, "ß": 30,
	}

	file, err := os.Open(PATH_TO_FILE)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	jumpDistances := []int{0, 0}
	var diff int

	for {

		jumpDistances = SmartAppend(jumpDistances, ReadRuneGetDistance(reader, germanAlphabetMap))

		jump := jumpDistances[1]
		if jumpDistances[0] < jumpDistances[1] {
			jump = jumpDistances[0]
		}

		_, err = reader.Discard(jump)
		if err != nil {
			log.Fatal(err)
		}

		jumpDistances = SmartAppend(jumpDistances, ReadRuneGetDistance(reader, germanAlphabetMap))

		diff = AbsInt(jumpDistances[0] - jumpDistances[1])
		_, err = reader.Discard(diff)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func ReadRuneGetDistance(reader *bufio.Reader, germanAlphabetMap map[string]int) int {
	readRune, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(string(readRune), " :Problem Reading: ", err)
	}

	readLetter := strings.ToLower(string(readRune))
	jA := germanAlphabetMap[readLetter]

	fmt.Println(readLetter, jA)
	return jA
}

func SmartAppend(array []int, element int) []int {
	if len(array) >= 2 {
		return []int{array[1], element}
	}
	return append(array, element)
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
