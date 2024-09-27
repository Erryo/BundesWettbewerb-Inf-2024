package main

import (
	"bufio"
	"fmt"
	"io"
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
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
		"ä": 27, "ö": 28, "ü": 29, "ß": 30,
	}
	numberOfMoves := make([]int, NUMBER_OF_PLAYERS)

	for playerIndex := range NUMBER_OF_PLAYERS {
		fmt.Println("-----------------------")
		// Gonna have to use Channels
		numberOfMoves[playerIndex] = getPlayersNoOfMoves(playerIndex, germanAlphabetMap)
		fmt.Println("-----------------------")
		fmt.Println(numberOfMoves)
	}
}

func getPlayersNoOfMoves(playerNo int, germanAlphabetMap map[string]int) int {
	file, err := os.Open(PATH_TO_FILE)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	playersMoves := 0

	// Not all players start from the same position
	// they are offset by their number
	_, err = reader.Discard(playerNo)
	if err != nil {
		log.Fatal(err)
	}

	for {
		jump := read(reader, germanAlphabetMap)
		fmt.Println(" ", jump)
		playersMoves += 1

		_, err := reader.Discard(jump)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err, jump)
			}
			break
		}
	}
	return playersMoves
}

// Reades the Current Rune under the cursor and increases the given pointer by 1
// if it is rune 195 (F*** umlaut BTW) it decreases the pointer and reads again
// until the read rune is a letter then returns the index of the Letter
func read(reader *bufio.Reader, germanAlphabetMap map[string]int) int {
	// better option, jump by the size of the read rune if it ain't size=1
	readRune, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}

	if readRune == 195 {
		fmt.Println("FUCKKKKK YOU 195")
		return read(reader, germanAlphabetMap)
	} else if !unicode.IsLetter(readRune) {
		return read(reader, germanAlphabetMap)
	}

	err = reader.UnreadRune()
	if err != nil {
		log.Fatal(err)
	}

	readLetter := strings.ToLower(string(readRune))
	fmt.Print(readLetter)
	return germanAlphabetMap[readLetter]
}
