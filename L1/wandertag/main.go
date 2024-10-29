package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// areaRange the start and end value of the portion
// availableTo is used to show which runner can particpate to this racePortion
// the index corresponds to the player index, True-can particpate
// False-Cannot particpate
type racePortion struct {
	start, end  int
	availableTo []bool
	noOfPlayers int
}

func main() {
	var raceMap []racePortion
	data := getData("../43.1/A3_Wandertag/wandern1.txt")

	firstPortion := racePortion{data[0][0], data[0][1], make([]bool, len(data)), 1}
	raceMap = append(raceMap, firstPortion)
	fmt.Println(raceMap)
	test1(data[1], 1, len(data), raceMap)
	// Overlapping
}

func test1(toAdd []int, index int, totalPlayers int, raceMap []racePortion) []racePortion {
	// index could be replaced with len(raceMap)-1
	c_min := toAdd[0]
	c_max := toAdd[1]
	for loopIndex, loopRacePortion := range raceMap {
		// Not Overlapping
		if c_min < loopRacePortion.start && c_max < loopRacePortion.start {
			// !!!!!!!! Problem
			// New portion not always at the end
			// !!!!!!!!!!!
			// Do not think it matters any more
			// i do not need it in order so i can just append
			// !!!!!!!!
			availableTo := make([]bool, totalPlayers)
			availableTo[index] = true
			newPortion := racePortion{start: toAdd[0], end: toAdd[1], availableTo: availableTo, noOfPlayers: 1}
			raceMap = append(raceMap, newPortion)

		} else if c_min > loopRacePortion.end && c_max > loopRacePortion.end {
			availableTo := make([]bool, totalPlayers)
			availableTo[index] = true
			newPortion := racePortion{start: toAdd[0], end: toAdd[1], availableTo: availableTo, noOfPlayers: 1}
			raceMap = append(raceMap, newPortion)
			// Co ranges
			// Right
		} else if c_min < loopRacePortion.end && c_min > loopRacePortion.start && c_max > loopRacePortion.end {
			toInsert := racePortion{start: c_min, end: loopRacePortion.end}
			newPortion := racePortion{start: loopRacePortion.end, end: c_max}
			loopRacePortion.end = c_min

			fmt.Println(loopRacePortion, toInsert, newPortion)
			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, toInsert, newPortion)
			// Left
		} else if c_max < loopRacePortion.end && c_max > loopRacePortion.start && c_min > loopRacePortion.end {
			toInsert := racePortion{start: loopRacePortion.start, end: c_max}
			newPortion := racePortion{start: c_min, end: loopRacePortion.start}
			loopRacePortion.start = c_max

			fmt.Println(loopRacePortion, toInsert, newPortion)
			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, toInsert, newPortion)
			// Sub-range
		} else if c_min > loopRacePortion.start && c_max < loopRacePortion.end {
			toInsert := racePortion{start: c_max, end: loopRacePortion.end}
			newPortion := racePortion{start: c_min, end: c_max}
			loopRacePortion.end = c_min

			fmt.Println(loopRacePortion, toInsert, newPortion)
			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, toInsert, newPortion)
			// Sup-range
		} else if c_max > loopRacePortion.start && c_min < loopRacePortion.end {
			toInsert := racePortion{start: c_min, end: loopRacePortion.start}
			newPortion := racePortion{start: loopRacePortion.end, end: c_max}

			fmt.Println(loopRacePortion, toInsert, newPortion)
			raceMap = append(raceMap, toInsert, newPortion)
		}
	}
	return raceMap
}

func insertToArray(index int, toInsert racePortion, raceMap []racePortion) []racePortion {
	return raceMap
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
