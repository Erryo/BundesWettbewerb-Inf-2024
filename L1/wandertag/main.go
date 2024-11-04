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
	for ind, date := range data {
		if ind == 0 {
			continue
		}
		fmt.Println("--------------------------------------------")
		fmt.Println(date)
		fmt.Println("IND", ind)
		raceMap = test1(date, ind, len(data), raceMap)
		fmt.Println("--------------------------------------------")
	}
	// Overlapping
	//for _, portion := range raceMap {
	//	fmt.Println(portion.start, portion.end, portion.noOfPlayers)
	//}
}

func test1(toAdd []int, index int, totalPlayers int, raceMap []racePortion) []racePortion {
	c_min := toAdd[0]
	c_max := toAdd[1]
	for _, portion := range raceMap {
		fmt.Println(portion.start, portion.end, portion.noOfPlayers)
	}
outer:
	for loopIndex, loopRacePortion := range raceMap {
		// Not Overlapping
		if c_min < loopRacePortion.start && c_max < loopRacePortion.start {
			fmt.Println("Not Overlapping left")
			// !!!!!!!! Problem
			// New portion not always at the end
			// !!!!!!!!!!!
			// Do not think it matters any more
			// i do not need it in order so i can just append
			// !!!!!!!!
			availableTo := make([]bool, totalPlayers)
			availableTo[index] = true
			for i, v := range raceMap {
				if i == index {
					continue
				}
				if v.start == toAdd[0] && v.end == toAdd[1] {
					continue outer
				}
			}
			newPortion := racePortion{start: toAdd[0], end: toAdd[1], availableTo: availableTo, noOfPlayers: 1}
			raceMap = append(raceMap, newPortion)

		} else if c_min > loopRacePortion.end && c_max > loopRacePortion.end {
			fmt.Println("Not Overlapping right")
			availableTo := make([]bool, totalPlayers)
			availableTo[index] = true

			for i, v := range raceMap {
				if i == index {
					continue
				}
				if v.start == toAdd[0] && v.end == toAdd[1] {
					continue outer
				}
			}
			newPortion := racePortion{start: toAdd[0], end: toAdd[1], availableTo: availableTo, noOfPlayers: 1}
			raceMap = append(raceMap, newPortion)
			// Co ranges
			// Right
		} else if c_min < loopRacePortion.end && c_min > loopRacePortion.start && c_max > loopRacePortion.end {

			fmt.Println("Co range right")
			toInsert := racePortion{start: c_min, end: loopRacePortion.end}
			newPortion := racePortion{start: loopRacePortion.end, end: c_max}
			loopRacePortion.end = c_min - 1

			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, toInsert, newPortion)
			// Left
		} else if c_max < loopRacePortion.end && c_max > loopRacePortion.start && c_min > loopRacePortion.end {
			fmt.Println("Co range left")
			fmt.Println(4)
			toInsert := racePortion{start: loopRacePortion.start, end: c_max}
			newPortion := racePortion{start: c_min, end: loopRacePortion.start}
			loopRacePortion.start = c_max

			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, toInsert, newPortion)
			// Sub-range
		} else if c_min > loopRacePortion.start && c_max < loopRacePortion.end {
			fmt.Println("Sub-range inside")
			toInsert := racePortion{start: c_max, end: loopRacePortion.end}
			newPortion := racePortion{start: c_min, end: c_max}
			loopRacePortion.end = c_min

			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, toInsert, newPortion)
			// Sup-range
		} else if c_max > loopRacePortion.end && c_min < loopRacePortion.start {
			fmt.Println("Sub-range outside")
			toInsert := racePortion{start: c_min, end: loopRacePortion.start}
			newPortion := racePortion{start: loopRacePortion.end, end: c_max}

			raceMap = append(raceMap, toInsert, newPortion)
		} else if c_min == loopRacePortion.start && c_max > loopRacePortion.end {
			fmt.Println("Equal Complete right")
			loopRacePortion.noOfPlayers += 1
			newPortion := racePortion{start: loopRacePortion.end + 1, end: c_max}

			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, newPortion)
		} else if c_min == loopRacePortion.end {
			fmt.Println("Equal Touch right")
			fmt.Println(7)
			loopRacePortion.end -= 1
			toInsert := racePortion{start: c_min, end: c_min}
			newPortion := racePortion{start: c_min + 1, end: c_max}

			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, newPortion)
			raceMap = append(raceMap, toInsert)
		} else if c_max == loopRacePortion.start {
			fmt.Println("Equal touch left")
			loopRacePortion.start += 1
			toInsert := racePortion{start: c_max, end: c_max}
			newPortion := racePortion{start: c_min, end: c_max - 1}

			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, newPortion)
			raceMap = append(raceMap, toInsert)
		} else if c_max == loopRacePortion.end && c_min < loopRacePortion.start {
			fmt.Println("Equal Complete left")
			newPortion := racePortion{start: c_min, end: loopRacePortion.end - 1}
			loopRacePortion.noOfPlayers += 1

			raceMap = append(raceMap, newPortion)
			raceMap[loopIndex] = loopRacePortion
		} else if c_min == loopRacePortion.start && c_max == loopRacePortion.end {
			fmt.Println("Equal equal")
			loopRacePortion.noOfPlayers += 1
			raceMap[loopIndex] = loopRacePortion
		} else if c_min == loopRacePortion.start && c_max < loopRacePortion.end {
			fmt.Println("Equal sub right")
			loopRacePortion.noOfPlayers += 1
			newPortion := racePortion{start: c_max + 1, end: loopRacePortion.end}
			loopRacePortion.start = c_max

			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, newPortion)
		} else if c_max == loopRacePortion.end && c_min > loopRacePortion.start {
			fmt.Println("Equal sub left")
			loopRacePortion.noOfPlayers += 1
			newPortion := racePortion{start: c_min + 1, end: loopRacePortion.end}
			loopRacePortion.end = c_min
			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, newPortion)
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

func cleanUp(raceMap []racePortion) []racePortion {
	for i, v := range raceMap {
	inner:
		for j, k := range raceMap {
			if i == j {
				continue
			}
			if v.start == k.start && v.end == k.end {
				raceMap = removeFromArr(raceMap, j)
				break inner
			}
		}
	}
	return raceMap
}

func removeFromArr(raceMap []racePortion, ind int) []racePortion {
	fmt.Println("Remove ", ind, raceMap[ind])
	left := raceMap[:ind]
	right := raceMap[ind+1:]
	return append(left, right...)
}
