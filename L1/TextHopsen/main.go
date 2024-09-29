package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

const PATH_TO_FILE string = "../43.1/J2_Texthopsen/test.txt"

const NUMBER_OF_PLAYERS = 2

// Indicates who should win in case of EOF error
// false - player 1
// true - player 2
var whoseJump bool

func main() {
	//	for i := range 6 {
	//		var PATH_TO_INPUT string = fmt.Sprintf("../43.1/J1_QuadratischPraktischGrБn/garten%v.txt", i)
	//	}

	var leastJumps uint8
	var noOfLeastJumps int
	data := readFileToArray()
	noOfLeastJumps = getPlayersNoOfMoves(0, &data)

	for playerIndex := range NUMBER_OF_PLAYERS {
		fmt.Println("-----------------------")
		noJumps := getPlayersNoOfMoves(uint8(playerIndex+1), &data)
		if noJumps < noOfLeastJumps {
			noOfLeastJumps = noJumps
			leastJumps = uint8(playerIndex) + 1
		}
		fmt.Println("-----------------------")

	}
	fmt.Printf("The winner is player number %v with %v jumps", leastJumps, noOfLeastJumps)
}

// Reads The File Rune by Rune appending the equevalent
// jumping values to an array
func readFileToArray() []uint8 {
	germanAlphabetMap := map[string]uint8{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8,
		"i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16,
		"q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23,
		"x": 24, "y": 25, "z": 26,
		"ä": 27, "ö": 28, "ü": 29, "ß": 30,
	}

	file, err := os.Open(PATH_TO_FILE)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer file.Close()

	var data []uint8
	reader := bufio.NewReader(file)

	jumpDistances := []int{0, 0}
	var diff int

	jumpDistances = SmartAppend(jumpDistances, ReadRuneGetDistance(reader, germanAlphabetMap)-2)
	jumpDistances = SmartAppend(jumpDistances, ReadRuneGetDistance(reader, germanAlphabetMap)-1)

	jump := jumpDistances[1]
	if jumpDistances[0] < jumpDistances[1] {
		jump = jumpDistances[0]
	}

	_, err = reader.Discard(jump)
	if err != nil {
		log.Fatal(err)
	}
	diff = AbsInt(jumpDistances[0] - jumpDistances[1])
	jumpDistances = SmartAppend(jumpDistances, ReadRuneGetDistance(reader, germanAlphabetMap)-1)

	if err = reader.UnreadRune(); err != nil {
		log.Fatal(err)
	}
	_, err = reader.Discard(diff)
	if err != nil {
		log.Fatal(err)
	}

	for range 3 {
		jumpDistances = SmartAppend(jumpDistances, ReadRuneGetDistance(reader, germanAlphabetMap))
		if err = reader.UnreadRune(); err != nil {
			log.Fatal(err)
		}

		jump := jumpDistances[1]
		if jumpDistances[0] < jumpDistances[1] {
			jump = jumpDistances[0]
		}

		_, err = reader.Discard(jump)
		if err != nil {
			log.Fatal(err)
		}
		diff = AbsInt(jumpDistances[0] - jumpDistances[1])

		jumpDistances = SmartAppend(jumpDistances, ReadRuneGetDistance(reader, germanAlphabetMap))
		if err = reader.UnreadRune(); err != nil {
			log.Fatal(err)
		}

		_, err = reader.Discard(diff)
		if err != nil {
			log.Fatal(err)
		}

	}
	return playersMoves
}

func increasePointer(locationPointer *int, amount uint8, lengthOfData int) error {
	if *locationPointer+int(amount) >= lengthOfData {
		return errors.New("pointer over the edge")
	}
	*locationPointer += int(amount)
	return nil
}

func ReadRuneGetDistance(reader *bufio.Reader, germanAlphabetMap map[string]int) int {
	readRune, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(string(readRune), " :Problem Reading: ", err)
	}

	if !unicode.IsLetter(readRune) {
		return ReadRuneGetDistance(reader, germanAlphabetMap)
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
