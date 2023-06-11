package word_search

func exist(board [][]byte, word string) bool {
	found := false
	max := len(word)
	m, n := len(board), len(board[0])

	var dfs func(r, c, i int)
	dfs = func(r, c, i int) {
		if found || i >= max || r < 0 || c < 0 || r >= m || c >= n || board[r][c] != word[i] {
			return
		}

		next := i + 1

		if next == max {
			found = true
			return
		}

		board[r][c] = '.'

		dfs(r+1, c, next)
		dfs(r-1, c, next)
		dfs(r, c+1, next)
		dfs(r, c-1, next)

		board[r][c] = word[i]
	}

	for row := 0; !found && row < m; row++ {
		for col := 0; !found && col < n; col++ {
			dfs(row, col, 0)
		}
	}

	return found
}
