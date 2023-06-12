package word_search_ii

const MAX int = 12

func findWords(board [][]byte, words []string) []string {
	trie := newTrie()
	var result []string

	seed(board, trie)

	for _, word := range words {
		if trie.find(word) {
			result = append(result, word)
		}
	}

	return result
}

// Unwralling all paths on board less than MAX moves long to the Trie, then search every word in Trie
func seed(board [][]byte, trie *Trie) {
	m, n := len(board), len(board[0])
	var current *TrieNode = trie.root

	var dfs func(r, c, i int)
	dfs = func(r, c, i int) {
		if i >= MAX || r < 0 || c < 0 || r >= m || c >= n || board[r][c] == '.' {
			return
		}

		next := i + 1

		if next == MAX {
			return
		}

		char := board[r][c]

		if current.links[char] == nil {
			current.links[char] = newTrieNode()
		}

		prev := current
		current = current.links[char]

		board[r][c] = '.'

		dfs(r+1, c, next)
		dfs(r-1, c, next)
		dfs(r, c+1, next)
		dfs(r, c-1, next)

		current = prev
		board[r][c] = char
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			dfs(row, col, 0)
			current = trie.root
		}
	}
}
