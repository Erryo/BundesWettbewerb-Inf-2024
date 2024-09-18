package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const PATH_TO_FILE string = "../43.1/J2_Texthopsen/hopsen1.txt"

func main() {
	// lookupTable := map[string]int{}
	//	for i := range 6 {
	//		var PATH_TO_INPUT string = fmt.Sprintf("../43.1/J1_QuadratischPraktischGrБn/garten%v.txt", i)
	//	}

	// Create a dict
	// Open file !!!NOT CLOSE IT!!!
	//
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

	for {
		readRune, size, err := reader.ReadRune()
		if err != nil {
			log.Fatal("Problem Reading")
		}
		if unicode.IsLetter(readRune) {
			err = reader.UnreadRune()
			if err != nil {
				log.Fatal("Error UnreadRune,", err)
			}
		}

		readLetter := strings.ToLower(string(readRune))
		jA := germanAlphabetMap[readLetter]
		fmt.Println(readLetter, readRune, jA, size)

		_, err = reader.Discard(jA)
		if err != nil {
			log.Fatal(err, jA, readLetter)
		}
	}
}
