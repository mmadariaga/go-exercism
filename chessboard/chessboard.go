package chessboard

import (
	"fmt"
)

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

func getValidFiles() []string {
	return []string{"A", "B", "C", "D", "E", "F", "G", "H"}
}

func isValidFile(file string) bool {

	validFiles := getValidFiles()

	for _, val := range validFiles {
		if val != file {
			continue
		}

		return true
	}

	return false

}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {

	if !isValidFile(file) {
		return 0
	}

	occupied := 0
	for _, val := range cb[file] {
		if val {
			occupied += 1
		}
	}

	return occupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {

	if rank > 8 || rank < 1 {
		return 0
	}

	occupied := 0
	for fileIdx, file := range cb {
		if !isValidFile(fileIdx) {
			fmt.Print("Not valid file " + fileIdx)
			return 0
		}

		if file[rank-1] {
			occupied += 1
		}
	}

	return occupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {

	sum := 0
	for idx := range cb {
		for range cb[idx] {
			sum += 1
		}
	}

	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	sum := 0
	for idx := range cb {
		for _, val := range cb[idx] {
			if val {
				sum += 1
			}
		}
	}

	return sum
}
