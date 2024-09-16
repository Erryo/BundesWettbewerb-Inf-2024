package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const PATH_TO_INPUT string = "../43.1/J1_QuadratischPraktischGrБn/garten0.txt"

func main() {
	readValues := readFile()
	heigth, width, noParties, maxNoLots := readValuesToInt(readValues)
	fmt.Println(heigth, width, noParties, maxNoLots)
}

func readFile() []string {
	file, err := os.Open(PATH_TO_INPUT)
	if err != nil {
		log.Fatal("Could not open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var readValues []string
	for scanner.Scan() {
		readValues = append(readValues, scanner.Text())
	}
	if len(readValues) < 3 || len(readValues) > 3 {
		log.Fatal("Len readValues incorrect")
	}
	return readValues
}

func readValuesToInt(readValues []string) (uint8, uint8, uint8, uint8) {
	var heigth uint8
	var width uint8
	var noParties uint8
	var maxNoLots uint8

	res, err := strconv.Atoi(readValues[0])
	if err != nil {
		log.Fatal("Please input valid numbers")
	}
	noParties = uint8(res)

	res, err = strconv.Atoi(readValues[1])
	if err != nil {
		log.Fatal("Please input valid numbers")
	}
	heigth = uint8(res)
	res, err = strconv.Atoi(readValues[2])
	if err != nil {
		log.Fatal("Please input valid numbers")
	}
	width = uint8(res)

	maxNoLots = noParties / 10
	return heigth, width, noParties, maxNoLots
}
