package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"unicode"
)

// const PATH_TO_FILE string = "../43.1/A2_Schwierigkeiten/schwierigkeiten0.txt"
func main() {
	for i := range 6 {
		var PATH_TO_FILE string = fmt.Sprintf("../43.1/A2_Schwierigkeiten/schwierigkeiten%v.txt", i)
		fmt.Println("===========================================")
		initalData, data, lettersToBe := readFile(PATH_TO_FILE)
		ratios := calcRatio(calcDistances(data, initalData[1]), initalData[1])
		fmt.Println(ratios)
		fmt.Println(orderRatios(lettersToBe, ratios))
	}
}

func orderRatios(lettersToBe []string, letterToRatio map[string]float32) []string {
	keys := make([]string, 0, len(letterToRatio))

	for key := range letterToRatio {
	inner:
		for _, letter := range lettersToBe {
			if key == letter {
				keys = append(keys, key)
				continue inner
			}
		}
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return letterToRatio[keys[i]] < letterToRatio[keys[j]]
	})
	return keys
}

// Calcutale the distance from the beggining to the letter (Input)
// and the distance from the letter to the end (Output)
func calcDistances(data [][]string, size int) map[string][]int {
	// var letterToDistances map[string][]int
	letterToDistances := make(map[string][]int, size)
	for _, exam := range data {
		for letter_id, letter := range exam {
			input := letter_id
			output := len(exam) - letter_id - 1

			letterArray := letterToDistances[letter]

			if len(letterArray) != 2 {
				letterArray = []int{0, 0}
			}
			letterArray[0] += input
			letterArray[1] += output
			letterToDistances[letter] = letterArray
		}
	}

	fmt.Println(letterToDistances)
	return letterToDistances
}

func calcRatio(letterToDistances map[string][]int, size int) map[string]float32 {
	letterToRatio := make(map[string]float32, size)
	for letter, distances := range letterToDistances {
		input_float := float32(distances[0])
		output_float := float32(distances[1])

		// Because a/0 is undefined i chnage it to 0.1 to give a better result
		// same thing for 0/a
		if input_float == 0 {
			input_float += 0.1
		}
		if output_float == 0 {
			output_float += 0.1
		}
		letterToRatio[letter] = input_float / output_float
	}
	return letterToRatio
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
	//_, _, err = reader.ReadRune()
	//if err != nil {
	//	log.Fatal(err)
	//}

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
	return startVariables, data, lettersToBe
}
