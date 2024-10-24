package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	getData("../43.1/A3_Wandertag/wandern1.txt")
}

func getData(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	noOfRunners := atoiNoErr(scanner.Text())
	runnerLimits := make([][]int, noOfRunners)

	ind := 0
	for scanner.Scan() {
		onesLimitString := strings.Split(scanner.Text(), " ")
		onesLimit := make([]int, 2)
		onesLimit[0] = atoiNoErr(onesLimitString[0])
		onesLimit[1] = atoiNoErr(onesLimitString[1])
		runnerLimits[ind] = append(runnerLimits[ind], onesLimit...)
		ind += 1

	}
	fmt.Println(runnerLimits)
	return runnerLimits
}

func atoiNoErr(num string) int {
	number, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return number
}
