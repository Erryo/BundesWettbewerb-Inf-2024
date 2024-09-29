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
	for {
		readRune, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if !unicode.IsLetter(readRune) {
			continue
		}
		readLetter := strings.ToLower(string(readRune))
		data = append(data, germanAlphabetMap[readLetter])
	}
	return data
}

func getPlayersNoOfMoves(playerNo uint8, data *[]uint8) int {
	playersMoves := 0
	var locationPointer int
	lengthOfData := len(*data)

	// Not all players start from the same position
	// they are offset by their number

	err := increasePointer(&locationPointer, playerNo, lengthOfData)
	if err != nil {
		return 0
	}

	for locationPointer < lengthOfData {

		fmt.Println("Location:", locationPointer)
		jump := (*data)[locationPointer]
		fmt.Println(" ", jump)
		playersMoves += 1

		err := increasePointer(&locationPointer, jump, lengthOfData)
		if err != nil {
			break
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
