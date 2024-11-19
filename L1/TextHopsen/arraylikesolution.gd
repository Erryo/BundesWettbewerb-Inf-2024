package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// Jump means to discard the smallest value
// Diff means to discard the difference
// between the bigger value and the jumped value
func Start() {
	germanAlphabetMap := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8,
		"i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16,
		"q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23,
		"x": 24, "y": 25, "z": 26,
		"ä": 27, "ö": 28, "ü": 29, "ß": 30,
	}

	file, err := os.Open(PATH_TO_FILE)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	pA, pB := 0, 0

	jA := read(reader, &pA, germanAlphabetMap)
	jB := read(reader, &pB, germanAlphabetMap)

	jumpDistance, firstPlayer := calcWhoseJump(jA, jB)
	differenceDistance := AbsInt(jA - jB)

	if firstPlayer {
		discard(reader, &pA, jumpDistance-2)
		jA = read(reader, &pA, germanAlphabetMap)
		discard(reader, &pB, differenceDistance)
		jB = read(reader, &pB, germanAlphabetMap)
	} else {
		// Subtracting 1 from jumpDistance because
		// the pointer is now at the letter after
		// the one From player 2/B
		discard(reader, &pB, jumpDistance-1)
		jB = read(reader, &pB, germanAlphabetMap)
		// this works because i am one letter further
		// than expected
		// smae as above
		discard(reader, &pA, differenceDistance)
		jA = read(reader, &pA, germanAlphabetMap)
	}
	unread(reader, &pA)

	// Start the Loop

	// An array that stores the values of the last 2 firstPlayer
	var lastJumps []bool = []bool{firstPlayer}
	for {
		jumpDistance, firstPlayer = calcWhoseJump(jA, jB)
		differenceDistance := AbsInt(jA - jB)

		lastJumps = SmartBoolAppend(lastJumps, firstPlayer)

		fmt.Println(jumpDistance, differenceDistance, jA, jB)
		fmt.Println(lastJumps)
		if firstPlayer {
			if lastJumps[0] == firstPlayer {
				discard(reader, &pA, jumpDistance)
			} else {
				discard(reader, &pA, jumpDistance-differenceDistance)
			}

			jA = read(reader, &pA, germanAlphabetMap)
			unread(reader, &pA)

			// differenceDistance
			discard(reader, &pB, differenceDistance)
			jB = read(reader, &pB, germanAlphabetMap)
		} else {
			if lastJumps[0] == firstPlayer {
				discard(reader, &pB, jumpDistance)
			} else {
				discard(reader, &pB, jumpDistance-differenceDistance)
			}

			jB = read(reader, &pB, germanAlphabetMap)
			unread(reader, &pB)

			// differenceDistance
			discard(reader, &pA, differenceDistance)
			jA = read(reader, &pA, germanAlphabetMap)
		}

		// If the pos of the current ToJump
		// is not the current pos
		// then i need to do
		// discardDistance = jumpDistance - diff
		// THIS ASSUMES UNREAD
	}
}

// Reades the Current Rune under the cursor and increases the given pointer by 1
// if it is rune 195  it decreases the pointer and reads again
// until the read rune is a letter then returns the index of the Letter
func read(reader *bufio.Reader, pX *int, germanAlphabetMap map[string]int) int {
	*pX += 1
	// better option, jump by the size of the read rune if it ain't size=1
	readRune, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	if readRune == 195 {
		*pX -= 1
		fmt.Println(" 195")
		return read(reader, pX, germanAlphabetMap)
	} else if !unicode.IsLetter(readRune) {
		return read(reader, pX, germanAlphabetMap)
	}
	readLetter := strings.ToLower(string(readRune))
	fmt.Println(readLetter)
	return germanAlphabetMap[readLetter]
}

func unread(reader *bufio.Reader, pX *int) {
	if err := reader.UnreadRune(); err != nil {
		log.Fatal(err)
	}
	*pX -= 1
}

func discard(reader *bufio.Reader, pX *int, distance int) {
	_, err := reader.Discard(distance)
	if err != nil {
		log.Fatal(distance, err)
	}
	*pX += distance
}

// takes in the jump distances of the 2 Players
// returns the smallest value and
// a bool - firstPlayer
// that indicates whether the first value(Coresponding
// to the first Player) is the one returned,
// and that the First Player was the last to jump
//
// Is Value B the bigger one
func calcWhoseJump(a, b int) (int, bool) {
	fmt.Println("calcWhoseJump:", a, b)
	if a < b {
		return a, true
	}
	return b, false
}

func SmartBoolAppend(array []bool, element bool) []bool {
	if len(array) >= 2 {
		return []bool{array[1], element}
	}
	return append(array, element)
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
