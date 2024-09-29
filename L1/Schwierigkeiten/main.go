package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

const PATH_TO_FILE string = "../43.1/A2_Schwierigkeiten/schwierigkeiten0.txt"

func main() {
	//for i := range 6 {
	//	var PATH_TO_INPUT string = fmt.Sprintf("../43.1/A2_Schwierigkeiten/schwierigkeiten%v.txt", i)
	//	readFile(PATH_TO_INPUT)
	//	fmt.Println("===========================================")
	//}
	_, data, lettersToBe := readFile(PATH_TO_FILE)
	fmt.Println(data, lettersToBe)
}

func readFile(path string) ([]int, [][]string, []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	// Read the first line
	var startVariables []int

outerLoop:
	for len(startVariables) < 3 {
		var number int
	innerLoop:
		for {
			readRune, _, err := reader.ReadRune()
			if err != nil {
				log.Fatal(err)
			}

			if readRune == 10 {
				startVariables = append(startVariables, number)
				break outerLoop
			}

			if !unicode.IsNumber(readRune) {
				break innerLoop
			}

			partialNumber, err := strconv.Atoi(string(readRune))
			if err != nil {
				log.Fatal(err)
			}
			number = (number * 10) + partialNumber
		}
		startVariables = append(startVariables, number)
	}

	// Read the n Lines of comparison
	var data [][]string
	_, _, err = reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}

	for range startVariables[0] {

		var lineData []string
		for {
			readRune, _, err := reader.ReadRune()
			if err != nil {
				log.Fatal(err)
			}

			if readRune == 10 {
				break
			}
			if !unicode.IsLetter(readRune) {
				continue
			}
			lineData = append(lineData, string(readRune))
		}
		data = append(data, lineData)
	}

	// Read the letters to be evaluated,ranked
	var lettersToBe []string
	for len(lettersToBe) < startVariables[2] {
		readRune, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if readRune == 10 {
			break
		}
		if !unicode.IsLetter(readRune) {
			continue
		}

		lettersToBe = append(lettersToBe, string(readRune))
	}
	fmt.Println(startVariables)
	fmt.Println(data)
	fmt.Println(lettersToBe)
	return startVariables, data, lettersToBe
}
