package word_search_ii

type cell struct {
	letter    byte
	isVisited bool
}

func outBounds(board *[][]cell, x, y int) bool {
	var height, widht = len(*board), len((*board)[0])

	return x >= widht || x < 0 || y >= height || y < 0
}

func scan(letter byte, board *[]string) *[][2]int {
	var result [][2]int

	for y, row := range *board {
		for x, l := range row {
			if byte(l) == letter {
				result = append(result, [2]int{x, y})
			}
		}
	}

	return &result
}

func convertBoard(board []string) [][]byte {
	result := make([][]byte, len(board))

	for i, word := range board {
		result[i] = make([]byte, len(word))
		for j, c := range word {
			result[i][j] = byte(c)
		}
	}

	return result
}