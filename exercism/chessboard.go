package main

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f := cb[file]
	var result int
	for _, occupied := range f {
		if occupied {
			result++
		}
	}
	return result
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	idx := rank - 1
	if idx < 0 || idx >= 8 {
		return 0
	}
	var result int
	for _, file := range cb {
		if file[idx] {
			result++
		}
	}
	return result
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * 8
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var result int
	for f := range cb {
		result += CountInFile(cb, f)
	}
	return result
}
