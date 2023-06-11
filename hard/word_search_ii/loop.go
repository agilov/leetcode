package word_search_ii

func findWordsLoop(board []string, words []string) []string {

	var result []string

	for _, word := range words {

		points := scan(word[0], &board)
		for _, start := range *points {
			b := setupBoard(&board)
			if search(word, b, start[0], start[1]) {
				result = append(result, word)
				break
			}
		}
	}

	return result
}

func search(word string, board *[][]cell, x, y int) bool {

	if word == "" {
		return true
	}

	if outBounds(board, x, y) {
		return false
	}

	visited := (*board)[y][x].isVisited
	letter := (*board)[y][x].letter

	if visited || letter != word[0] {
		return false
	}

	(*board)[y][x].isVisited = true

	variants := [][2]int{
		[2]int{x + 1, y},
		[2]int{x - 1, y},
		[2]int{x, y + 1},
		[2]int{x, y - 1},
	}

	for _, variant := range variants {
		if search(word[1:], board, variant[0], variant[1]) {
			return true
		}
	}
	(*board)[y][x].isVisited = false
	return false
}

func setupBoard(board *[]string) *[][]cell {
	result := make([][]cell, len(*board))

	for i, line := range *board {
		result[i] = make([]cell, len(line))

		for j, letter := range line {
			result[i][j] = cell{letter: byte(letter), isVisited: false}
		}
	}

	return &result
}
