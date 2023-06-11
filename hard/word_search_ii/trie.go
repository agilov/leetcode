package word_search_ii

const MAX int = 12

type TrieNode struct {
	links map[byte]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) add(word string) {
	current := t.root

	for _, c := range word {
		if current.links[byte(c)] == nil {
			current.links[byte(c)] = newTrieNode()
		}
		current = current.links[byte(c)]
	}
}

func (t *Trie) find(word string) bool {
	current := t.root

	for i := 0; i < len(word); i++ {
		char := word[i]
		if current.links[char] == nil {
			return false
		}
		current = current.links[word[i]]
	}

	return true
}

func newTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

func newTrieNode() *TrieNode {
	return &TrieNode{links: make(map[byte]*TrieNode)}
}

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
