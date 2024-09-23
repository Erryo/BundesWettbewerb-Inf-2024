package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func Start() {
	germanAlphabetMap := map[string]int{
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
	reader := bufio.NewReader(file)

	pA, pB := 0, 0
	fmt.Println(pA, pB)
	read(reader, &pA, germanAlphabetMap)
	fmt.Println(pA, pB)
}

// Reades the Current Rune under the cursor and increases the given pointer by 1
// if it is rune 195 (F*** umlaut BTW) it decreases the pointer and reads again
// until the read rune is a letter then returns the index of the Letter
func read(reader *bufio.Reader, pX *int, germanAlphabetMap map[string]int) int {
	*pX += 1
	readRune, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	// 182 164 171 159
	if readRune == 195 {
		*pX -= 1
		fmt.Println("FUCKKKKK YOU 195")
		return read(reader, pX, germanAlphabetMap)
	} else if !unicode.IsLetter(readRune) {
		return read(reader, pX, germanAlphabetMap)
	}
	return germanAlphabetMap[string(readRune)]
}
