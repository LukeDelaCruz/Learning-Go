package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

var files = [...]string{"A", "B", "C", "D", "E", "F", "G", "H"}

// CountInFile returns how many squares are occupied in the chessboard, within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0

	for _, rank := range cb[file] {
		if rank {
			count++
		}
	}

	return count
}

// CountInRank returns how many squares are occupied in the chessboard, within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0

	if rank < 1 || rank > 8 {
		return 0
	}

	for _, file := range files {
		if cb[file][rank-1] {
			count++
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0

	for range files {
		count += 8
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0

	for _, file := range files {
		for rank := 0; rank < 8; rank++ {
			if cb[file][rank] {
				count++
			}
		}
	}

	return count
}
