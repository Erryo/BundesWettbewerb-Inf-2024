package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const PATH_TO_FILE string = "../43.1/J2_Texthopsen/hopsen1.txt"

const NUMBER_OF_PLAYERS = 2

func main() {
	// lookupTable := map[string]int{}
	//	for i := range 6 {
	//		var PATH_TO_INPUT string = fmt.Sprintf("../43.1/J1_QuadratischPraktischGrБn/garten%v.txt", i)
	//	}

	germanAlphabetMap := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8,
		"i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16,
		"q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23,
		"x": 24, "y": 25, "z": 26,
		"ä": 27, "ö": 28, "ü": 29, "ß": 30,
	}
	numberOfMoves := make([]int, NUMBER_OF_PLAYERS)

	data := readFileToString()
	for playerIndex := range NUMBER_OF_PLAYERS {
		fmt.Println("-----------------------")
		// Gonna have to use Channels
		numberOfMoves[playerIndex] = getPlayersNoOfMoves(playerIndex, data, germanAlphabetMap)
		fmt.Println("-----------------------")
		fmt.Println(numberOfMoves)
	}
}

func readFileToString() string {
	file, err := os.Open(PATH_TO_FILE)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer file.Close()

	var data string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data += scanner.Text()
	}
	return data
}

func getPlayersNoOfMoves(playerNo int, data string, germanAlphabetMap map[string]int) int {
	playersMoves := 0
	var locationPointer int
	lengthOfData := len(data)

	// Not all players start from the same position
	// they are offset by their number

	err := increasePointer(&locationPointer, playerNo, lengthOfData)
	if err != nil {
		return 0
	}

	for locationPointer < lengthOfData {
		jump := read(data, &locationPointer, germanAlphabetMap)
		fmt.Println(" ", jump)
		playersMoves += 1

		err := increasePointer(&locationPointer, jump, lengthOfData)
		if err != nil {
			break
		}
	}
	return playersMoves
}

func increasePointer(locationPointer *int, amount, lengthOfData int) error {
	if *locationPointer+amount >= lengthOfData {
		return errors.New("pointer over the edge")
	}
	*locationPointer += amount
	return nil
}

// Reades the Current Rune under the cursor and increases the given pointer by 1
// if it is rune 195 (F*** umlaut BTW) it decreases the pointer and reads again
// until the read rune is a letter then returns the index of the Letter
func read(data string, locationPointer *int, germanAlphabetMap map[string]int) int {
	// better option, jump by the size of the read rune if it ain't size=1
	readRune := data[*locationPointer]
	*locationPointer += 1

	if readRune == 195 {
		fmt.Println("FUCKKKKK YOU 195")
		return read(data, locationPointer, germanAlphabetMap)
	} else if !unicode.IsLetter(rune(readRune)) {
		fmt.Println("NOT A LETTER", string(readRune))
		return read(data, locationPointer, germanAlphabetMap)
	}

	*locationPointer -= 1

	readLetter := strings.ToLower(string(readRune))
	fmt.Print(readLetter)
	return germanAlphabetMap[readLetter]
}
