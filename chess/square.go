package chess

import "fmt"

type Square struct {
	File string
	Rank int
}

func IndexToSquare(index int) Square {
	// Check if the index is within the playable area (26 to 117)
	if index < 26 || index > 117 {
		return Square{} // Non-playable area
	}

	// Calculate the offset from the playable starting square (a8)
	offset := index - 26

	// Calculate the file (column) and rank (row) for the 8x8 playable grid
	file := offset % 8       // 0-7 for file (a-h)
	rank := 7 - (offset / 8) // 7-0 for rank (8-1)

	// Convert the file index (0-7) to letters (a-h)
	fileLetter := string(rune('a' + file))

	// Convert the rank index (0-7) to ranks (8-1)
	rankInt := rank + 1
	fmt.Printf(" index: %d , file: %s, rank: %d\n", index, fileLetter, rankInt)

	// Return the square with calculated file and rank
	return Square{
		File: fileLetter,
		Rank: rankInt,
	}
}
