package main

import (
	"bufio"
	"io"
	"log"
	"os"
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
		"ä": 27, "ö": 28, "ü": 29,
		"A": 30, "B": 31, "C": 32, "D": 33, "E": 34, "F": 35, "G": 36, "H": 37, "I": 38, "J": 39, "K": 40, "L": 41, "M": 42, "N": 43, "O": 44, "P": 45, "Q": 46, "R": 47, "S": 48, "T": 49, "U": 50, "V": 51, "W": 52, "X": 53, "Y": 54, "Z": 55,
		"Ä": 56, "Ö": 57, "Ü": 58,
	}

	file, err := os.Open(PATH_TO_FILE)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		readWord, err := reader.ReadBytes(32)
		if err != nil && err != io.EOF {
			log.Fatal("Problem Reading")
		}
		calculateJump(germanAlphabetMap, string(readWord))
	}
}

func calculateJump(germanAlphabetMap map[string]int, readWord string) {
}
