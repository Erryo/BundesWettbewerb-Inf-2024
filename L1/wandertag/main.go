package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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
type ByNoOfPlayers []racePortion

func (a ByNoOfPlayers) Len() int           { return len(a) }
func (a ByNoOfPlayers) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNoOfPlayers) Less(i, j int) bool { return a[i].noOfPlayers > a[j].noOfPlayers }

func main() {
	startTime := time.Now()
	var raceMap []racePortion
	data := getData("../43.1/A3_Wandertag/wandern1.txt")

	firstPortion := racePortion{data[0][0], data[0][1], make([]bool, len(data)), 1}
	raceMap = append(raceMap, firstPortion)
	fmt.Println(raceMap)
	for ind, datum := range data {
		if ind == 0 {
			continue
		}
		fmt.Println("------------------", ind, datum, "--------------------------")
		raceMap = test1(datum, ind, len(data), raceMap)
		fmt.Println("--------------------------------------------")
	}
	// Overlapping

	printMap(raceMap)

	getHighest(raceMap)
	fmt.Println(time.Since(startTime))
}

func test1(toAdd []int, index int, totalPlayers int, raceMap []racePortion) []racePortion {
	c_min := toAdd[0]
	c_max := toAdd[1]

outer:
	for loopIndex, loopRacePortion := range raceMap {
		// Not Overlapping
		if c_min < loopRacePortion.start && c_max < loopRacePortion.start {

			fmt.Println("Not Overlapping left")
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

			availableTo := make([]bool, totalPlayers)
			availableTo[index] = true
			availableTo[loopIndex] = true
			newAvailableTo := make([]bool, totalPlayers)
			newAvailableTo[index] = true

			toInsert := racePortion{start: c_min, end: loopRacePortion.end, availableTo: availableTo, noOfPlayers: loopRacePortion.noOfPlayers + 1}
			newPortion := racePortion{start: loopRacePortion.end, end: c_max, noOfPlayers: 1, availableTo: newAvailableTo}
			loopRacePortion.end = c_min - 1

			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, toInsert)
			if index < 2 {
				raceMap = append(raceMap, newPortion)
			} else {
				ix := findIndex(raceMap, c_min, c_max)
				raceMap[ix].availableTo[loopIndex] = true
				raceMap[ix].noOfPlayers += 1
			}
			// Left
		} else if c_max < loopRacePortion.end && c_max > loopRacePortion.start && c_min > loopRacePortion.end {
			fmt.Println("Co range left")
			availableTo := make([]bool, totalPlayers)
			availableTo[index] = true
			availableTo[loopIndex] = true
			newAvailableTo := make([]bool, totalPlayers)
			newAvailableTo[index] = true

			toInsert := racePortion{start: loopRacePortion.start, end: c_max, availableTo: availableTo, noOfPlayers: loopRacePortion.noOfPlayers + 1}
			newPortion := racePortion{start: c_min, end: loopRacePortion.start, availableTo: newAvailableTo, noOfPlayers: 1}
			loopRacePortion.start = c_max

			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, toInsert)
			if index < 2 {
				raceMap = append(raceMap, newPortion)
			} else {
				ix := findIndex(raceMap, c_min, c_max)
				raceMap[ix].availableTo[loopIndex] = true
				raceMap[ix].noOfPlayers += 1
			}
			// Sub-range
		} else if c_min > loopRacePortion.start && c_max < loopRacePortion.end {
			fmt.Println("Sub-range inside")
			availableTo := make([]bool, totalPlayers)
			availableTo[index] = true
			availableTo[loopIndex] = true
			newAvailableTo := make([]bool, totalPlayers)
			newAvailableTo[index] = true

			toInsert := racePortion{start: c_max, end: loopRacePortion.end, availableTo: newAvailableTo, noOfPlayers: loopRacePortion.noOfPlayers}
			newPortion := racePortion{start: c_min, end: c_max, availableTo: availableTo, noOfPlayers: loopRacePortion.noOfPlayers + 1}
			loopRacePortion.end = c_min // a

			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, toInsert)
			if index < 2 {
				raceMap = append(raceMap, newPortion)
			} else {
				ix := findIndex(raceMap, c_min, c_max)
				raceMap[ix].availableTo[loopIndex] = true
				raceMap[ix].noOfPlayers += 1
			}
			// Sup-range
		} else if c_max > loopRacePortion.end && c_min < loopRacePortion.start {
			fmt.Println("Sub-range outside")
			availableTo := make([]bool, totalPlayers)
			availableTo[index] = true

			toInsert := racePortion{start: c_min, end: loopRacePortion.start, availableTo: availableTo, noOfPlayers: 1}
			newPortion := racePortion{start: loopRacePortion.end, end: c_max, availableTo: availableTo, noOfPlayers: 1}
			loopRacePortion.availableTo[index] = true
			loopRacePortion.noOfPlayers += 1

			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, toInsert)
			if index < 2 {
				raceMap = append(raceMap, newPortion)
			} else {
				ix := findIndex(raceMap, c_min, c_max)
				raceMap[ix].noOfPlayers += 1
			}
		} else if c_min == loopRacePortion.start && c_max > loopRacePortion.end {
			fmt.Println("Equal Complete right")

			availableTo := make([]bool, totalPlayers)
			availableTo[index] = true

			loopRacePortion.noOfPlayers += 1
			loopRacePortion.availableTo[index] = true
			newPortion := racePortion{start: loopRacePortion.end + 1, end: c_max, availableTo: availableTo, noOfPlayers: 1}

			raceMap[loopIndex] = loopRacePortion
			if index < 2 {
				raceMap = append(raceMap, newPortion)
			} else {
				ix := findIndex(raceMap, c_min, c_max)
				raceMap[ix].noOfPlayers += 1
			}
		} else if c_min == loopRacePortion.end {
			fmt.Println("Equal Touch right")

			availableTo := loopRacePortion.availableTo
			availableTo[index] = true
			newAvailableTo := make([]bool, totalPlayers)
			newAvailableTo[index] = true

			loopRacePortion.end -= 1
			toInsert := racePortion{start: c_min, end: c_min, availableTo: availableTo, noOfPlayers: loopRacePortion.noOfPlayers + 1}
			newPortion := racePortion{start: c_min + 1, end: c_max, availableTo: newAvailableTo, noOfPlayers: 1}

			raceMap[loopIndex] = loopRacePortion
			if index < 2 {
				raceMap = append(raceMap, newPortion)
			} else {
				ix := findIndex(raceMap, c_min, c_max)
				raceMap[ix].noOfPlayers += 1
			}
			raceMap = append(raceMap, toInsert)
		} else if c_max == loopRacePortion.start {
			fmt.Println("Equal touch left")

			availableTo := loopRacePortion.availableTo
			availableTo[index] = true
			newAvailableTo := make([]bool, totalPlayers)
			newAvailableTo[index] = true

			loopRacePortion.start += 1
			toInsert := racePortion{start: c_max, end: c_max, availableTo: availableTo, noOfPlayers: loopRacePortion.noOfPlayers + 1}
			newPortion := racePortion{start: c_min, end: c_max - 1, availableTo: newAvailableTo, noOfPlayers: 1}

			raceMap[loopIndex] = loopRacePortion
			if index < 2 {
				raceMap = append(raceMap, newPortion)
			} else {
				ix := findIndex(raceMap, c_min, c_max)
				raceMap[ix].noOfPlayers += 1
			}
			raceMap = append(raceMap, toInsert)
		} else if c_max == loopRacePortion.end && c_min < loopRacePortion.start {
			fmt.Println("Equal Complete left")

			availableTo := make([]bool, totalPlayers)
			availableTo[index] = true

			newPortion := racePortion{start: c_min, end: loopRacePortion.end - 1, availableTo: availableTo, noOfPlayers: 1}
			loopRacePortion.noOfPlayers += 1

			if index < 2 {
				raceMap = append(raceMap, newPortion)
			} else {
				ix := findIndex(raceMap, c_min, c_max)
				raceMap[ix].noOfPlayers += 1
			}
			raceMap[loopIndex] = loopRacePortion
		} else if c_min == loopRacePortion.start && c_max == loopRacePortion.end {
			fmt.Println("Equal equal")

			loopRacePortion.noOfPlayers += 1
			loopRacePortion.availableTo[index] = true
			raceMap[loopIndex] = loopRacePortion
		} else if c_min == loopRacePortion.start && c_max < loopRacePortion.end {
			fmt.Println("Equal sub right")

			availableTo := make([]bool, totalPlayers)
			availableTo[index] = true

			loopRacePortion.noOfPlayers += 1
			loopRacePortion.availableTo[index] = true
			newPortion := racePortion{start: c_max + 1, end: loopRacePortion.end, availableTo: availableTo, noOfPlayers: 1}
			loopRacePortion.start = c_max

			raceMap[loopIndex] = loopRacePortion
			if index < 2 {
				raceMap = append(raceMap, newPortion)
			} else {
				ix := findIndex(raceMap, c_min, c_max)
				raceMap[ix].noOfPlayers += 1
			}
		} else if c_max == loopRacePortion.end && c_min > loopRacePortion.start {
			fmt.Println("Equal sub left")

			availableTo := make([]bool, totalPlayers)
			availableTo[index] = true

			loopRacePortion.noOfPlayers += 1
			newPortion := racePortion{start: c_min + 1, end: loopRacePortion.end, availableTo: availableTo, noOfPlayers: 1}
			loopRacePortion.end = c_min
			loopRacePortion.availableTo[index] = true

			raceMap[loopIndex] = loopRacePortion
			if index < 2 {
				raceMap = append(raceMap, newPortion)
			} else {
				ix := findIndex(raceMap, c_min, c_max)
				raceMap[ix].noOfPlayers += 1
			}
		}
	}
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

func getHighest(raceMap []racePortion) {
	selected := []racePortion{}
	for i := range 3 {

		sort.Sort(ByNoOfPlayers(raceMap))
		selected = append(selected, raceMap[0])
		raceMap = remove(raceMap, 0)
		raceMap = recalcNoOfPlayers(selected[i].availableTo, raceMap)
	}
	printMap(raceMap)
	printMap(selected)
}

func atoiNoErr(num string) int {
	number, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func findIndex(raceMap []racePortion, start, end int) int {
	for i, v := range raceMap {
		if v.start == start && v.end == end {
			return i
		}
	}
	return -1
}

func remove(slice []racePortion, s int) []racePortion {
	return append(slice[:s], slice[s+1:]...)
}

func printMap(raceMap []racePortion) {
	fmt.Println("--------------------------------------------")
	for _, portion := range raceMap {
		fmt.Println(portion.start, portion.end, portion.noOfPlayers, portion.availableTo)
	}
	fmt.Println("--------------------------------------------")
}

func recalcNoOfPlayers(fromRemoved []bool, raceMap []racePortion) []racePortion {
	for _, portion := range raceMap {
		for i, value := range portion.availableTo {
			if fromRemoved[i] == true && value {
				portion.noOfPlayers -= 1
				portion.availableTo[i] = false
			}
		}
	}
	return raceMap
}
