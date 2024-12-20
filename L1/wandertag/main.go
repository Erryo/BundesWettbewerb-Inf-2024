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

	availableTo := make([]bool, len(data))
	availableTo[0] = true
	firstPortion := racePortion{data[0][0], data[0][1], availableTo, 0}

	raceMap = append(raceMap, firstPortion)
	fmt.Println(raceMap)

	totalPlayers := len(data)
	for ind, datum := range data {
		if ind == 0 {
			continue
		}
		fmt.Println("------------------", ind, datum, "--------------------------")
		// raceMap = test1(datum, ind, len(data), raceMap)
		raceMap = addNew(raceMap, ind, totalPlayers, datum[0], datum[1])
		// fmt.Println("--------------------------------------------")
	}
	// Overlapping

	printMap(raceMap)

	calcNumPlayers(raceMap)
	printMap(raceMap)

	getHighest(raceMap)

	fmt.Println(time.Since(startTime))
}

func addNew(raceMap []racePortion, runnerIndex, totalPlayers, cMin, cMax int) []racePortion {
	var added bool
	availableTo := make([]bool, totalPlayers)
	availableTo[runnerIndex] = true
	for loopIndex, loopRacePortion := range raceMap {
		if cMin < loopRacePortion.start && cMax < loopRacePortion.start {
			fmt.Println("Not Overlapping left")
			if added {
				continue
			}
			fPortion := racePortion{start: cMin, end: cMax, availableTo: availableTo}
			added = true
			raceMap = append(raceMap, fPortion)

		} else if cMin == loopRacePortion.start && cMax == loopRacePortion.end {
			fmt.Println("Equal equal")
		} else if cMin > loopRacePortion.end && cMax > loopRacePortion.end {
			fmt.Println("Not Overlapping right")
			if added {
				continue
			}
			fPortion := racePortion{start: cMin, end: cMax, availableTo: availableTo}
			added = true
			raceMap = append(raceMap, fPortion)
		} else if cMin < loopRacePortion.end && cMin > loopRacePortion.start && cMax > loopRacePortion.end {
			fmt.Println("Co range right")

			fAvailableTo := make([]bool, totalPlayers)
			copy(fAvailableTo, loopRacePortion.availableTo)
			fAvailableTo[runnerIndex] = true

			fPortion := racePortion{start: cMin, end: loopRacePortion.end, availableTo: fAvailableTo}
			// sPortion := racePortion{start: loopRacePortion.end + 1, end: cMax, availableTo: availableTo}

			raceMap = addNew(raceMap, runnerIndex, totalPlayers, loopRacePortion.end+1, cMax)
			loopRacePortion.end = cMin - 1

			raceMap[loopIndex] = loopRacePortion
			raceMap = append(raceMap, fPortion)
			// raceMap = append(raceMap, sPortion)
		} else if cMax < loopRacePortion.end && cMax > loopRacePortion.start && cMin < loopRacePortion.start {
			fmt.Println("Co range left")

			fAvailableTo := make([]bool, totalPlayers)
			copy(fAvailableTo, loopRacePortion.availableTo)
			fAvailableTo[runnerIndex] = true

			fPortion := racePortion{start: loopRacePortion.start, end: cMax, availableTo: fAvailableTo}
			raceMap = addNew(raceMap, runnerIndex, totalPlayers, cMin, loopRacePortion.start-1)
			loopRacePortion.end = cMax
			raceMap = append(raceMap, fPortion)

		} else if cMin > loopRacePortion.start && cMax < loopRacePortion.end {
			fmt.Println("Sub-range inside")

			fAvailableTo := make([]bool, totalPlayers)
			copy(fAvailableTo, loopRacePortion.availableTo)
			fAvailableTo[runnerIndex] = true

			// Not using recursion because i am inserting, not creating a new portion
			fPortion := racePortion{start: cMin, end: cMax, availableTo: fAvailableTo}
			raceMap = append(raceMap, fPortion)
			sPortion := racePortion{start: cMax + 1, end: loopRacePortion.end, availableTo: availableTo}
			loopRacePortion.end = cMin - 1

			raceMap = append(raceMap, sPortion)
			raceMap[loopIndex] = loopRacePortion

		} else if cMax > loopRacePortion.end && cMin < loopRacePortion.start {
			fmt.Println("Sub-range outside")
		} else if cMin == loopRacePortion.start && cMax > loopRacePortion.end {
			fmt.Println("Equal Complete right")
		} else if cMin == loopRacePortion.end {
			fmt.Println("Equal Touch right")
		} else if cMax == loopRacePortion.start {
			fmt.Println("Equal touch left")
		} else if cMax == loopRacePortion.end && cMin < loopRacePortion.start {
			fmt.Println("Equal Complete left")
		} else if cMin == loopRacePortion.start && cMax < loopRacePortion.end {
			fmt.Println("Equal sub right")
		} else if cMax == loopRacePortion.end && cMin > loopRacePortion.start {
			fmt.Println("Equal sub left")
		}
		printMap(raceMap)
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
	// printMap(selected)
	declareRaces(selected)
}

func declareRaces(selected []racePortion) {
	for k, portion := range selected {
		fmt.Printf("%v.Portion: %v \n", k, (portion.start+portion.end)/2)
		for i, v := range portion.availableTo {
			if v {
				fmt.Print(i+1, " ")
			}
		}
		fmt.Print("\n")
	}
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
	fmt.Println("[[[o[o[o[o[[o")
	// printMap(raceMap)
	for j := 0; j < len(raceMap)-1; {
		portion := raceMap[j]
		fmt.Println(portion)
	inner:
		for i, value := range portion.availableTo {
			fmt.Println("...", i, "...")
			if fromRemoved[i] && value {
				if portion.noOfPlayers-1 > 0 {
					fmt.Println("<>")
					portion.availableTo[i] = false
					portion.noOfPlayers -= 1
					raceMap[j] = portion
				} else {
					fmt.Println("><")
					raceMap = remove(raceMap, j)
					j--
					break inner
				}
			}
		}
		j++
		fmt.Println(portion)
	}
	// printMap(raceMap)
	return raceMap
}

func calcNumPlayers(raceMap []racePortion) []racePortion {
	for i, portion := range raceMap {
		number := 0
		for _, datum := range portion.availableTo {
			if datum == true {
				number++
			}
		}
		portion.noOfPlayers = number
		raceMap[i] = portion
	}
	return raceMap
}

// AN or operation
func boolOr(a, b []bool) []bool {
	for i := 0; i < len(a); i++ {
		if b[i] {
			a[i] = true
		}
	}
	return a
}
