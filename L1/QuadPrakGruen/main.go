package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const PATH_TO_INPUT string = "../43.1/J1_QuadratischPraktischGrБn/garten5.txt"

func main() {
	readValues := readFile()
	heigth, width, noParties, _ := readValuesToInt(readValues)

	totalArea := heigth * width
	areaPerPartie := float64(totalArea) / float64(noParties)
	lengthSquare := math.Sqrt(areaPerPartie)
	divA := float64(heigth) / lengthSquare
	divB := float64(width) / lengthSquare

	text := fmt.Sprintf("For a lot of the size %vx%v and %v parties Mr.Green should give each person %.4f sqr meters \n", heigth, width, noParties, areaPerPartie)
	text_2 := fmt.Sprintf("which means he should divide the heigth by %.4f and the width by %.4f\n", divA, divB)
	fmt.Println(text, text_2)
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

func readValuesToInt(readValues []string) (int, int, int, int) {
	var heigth int
	var width int
	var noParties int
	var maxNoLots int

	res, err := strconv.Atoi(readValues[0])
	if err != nil {
		log.Fatal("Please input valid numbers")
	}
	noParties = res

	res, err = strconv.Atoi(readValues[1])
	if err != nil {
		log.Fatal("Please input valid numbers")
	}
	heigth = res
	res, err = strconv.Atoi(readValues[2])
	if err != nil {
		log.Fatal("Please input valid numbers")
	}
	width = res

	maxNoLots = noParties / 10
	return heigth, width, noParties, maxNoLots
}
