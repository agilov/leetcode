package word_search_ii

type Node struct {
	isWord     bool
	word       string
	liknsCount uint8
	links      [26]*Node
}

// Finding words in board using trie approach + DFS with removing words from trie after word is found
func findWords(board [][]byte, words []string) []string {
	current := buildTrie(words)
	result := make([]string, 0)
	m, n := len(board), len(board[0])

	// closure dfs function used to improve code readability
	// and to keep static variables and backtracking pointers in closure
	var dfs func(r, c, i int)
	dfs = func(r, c, i int) {
		// checking boundaries and if board cell is visited. Also checking character is present in the trie
		if r < 0 || c < 0 || r == m || c == n || board[r][c] == 0 || current.links[board[r][c]-'a'] == nil {
			return
		}

		char := board[r][c]
		charidx := char - 'a'

		// if word is found - then we remove word mark from this node and check if it's leaf node
		if current.links[charidx].isWord {
			result = append(result, current.links[charidx].word)
			current.links[charidx].isWord = false

			// if this word is leaf node (there is no further words with this prefix)
			// this node should be removed from the tree
			if current.links[charidx].liknsCount == 0 {
				current.links[charidx] = nil
				current.liknsCount--
				return
			}
		}

		// Setup backtrack variables
		prev := current
		next := i + 1
		board[r][c] = 0
		current = current.links[charidx]

		// search in neighbour nodes with dfs
		dfs(r+1, c, next)
		dfs(r-1, c, next)
		dfs(r, c+1, next)
		dfs(r, c-1, next)

		// backtracking
		current = prev
		board[r][c] = char

		// removing hanging leaves that are not words and has no more links
		if !current.links[charidx].isWord && current.links[charidx].liknsCount == 0 {
			current.links[charidx] = nil
			current.liknsCount--
		}
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			// skip characters not present in trie
			if current.links[board[row][col]-'a'] != nil {
				dfs(row, col, 0)
			}
		}
	}

	return result
}

// Adding all words to trie and returning it
// We're building entire trie in this function to optimize runtime speed
func buildTrie(words []string) *Node {
	trie := &Node{}

	for _, word := range words {

		current := trie

		for i := 0; i < len(word); i++ {
			idx := word[i] - 'a'
			if current.links[idx] == nil {
				current.links[idx] = &Node{}
				current.liknsCount++
			}
			current = current.links[idx]
		}

		current.isWord = true
		current.word = word
	}

	return trie
}
