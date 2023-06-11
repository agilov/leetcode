package word_search_ii

type TrieNode struct {
	links [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) add(word string) {
	current := t.root

	for _, c := range word {
		index := c - 'a'
		if current.links[index] == nil {
			current.links[index] = &TrieNode{}
		}
		current = current.links[index]
	}
}

func (t *Trie) find(word string) bool {
	current := t.root

	for _, c := range word {
		if current.links[c-'a'] == nil {
			return false
		}
		current = current.links[c-'a']
	}

	return true
}

func newTrie() *Trie {
	return &Trie{root: &TrieNode{}}
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

func findWords(board []string, words []string) []string {
	trie := newTrie()
	brd := convertBoard(board)

	var result []string

	for y := 0; y < len(brd); y++ {
		for x := 0; x < len(brd[0]); x++ {
			insert(copyBoard(brd), "", trie, x, y, 0)
		}
	}

	for _, word := range words {
		if trie.find(word) {
			result = append(result, word)
		}
	}

	return result
}

func copyBoard(board [][]byte) [][]byte {
	result := make([][]byte, len(board))
	for i := range board {
		result[i] = make([]byte, len(board[i]))
		copy(result[i], board[i])
	}
	return result
}

// Inserting board to trie from concrete letter.
// Each letter have at most 4 neighbours, if neigbour is not inserted before - it marged as inserted after insertion
// insertion stops when there is no not-inserted letters left
func insert(board [][]byte, prefix string, trie *Trie, x int, y int, level int) {

	if outOfBounds(&board, x, y) || level > 10 {
		return
	}

	if board[y][x] == 255 {
		return
	}

	toInsert := prefix + string(board[y][x])

	trie.add(toInsert)

	board[y][x] = 255

	variants := [][2]int{
		[2]int{x + 1, y},
		[2]int{x - 1, y},
		[2]int{x, y + 1},
		[2]int{x, y - 1},
	}

	for _, variant := range variants {
		insert(copyBoard(board), toInsert, trie, variant[0], variant[1], level + 1)
	}
}

func outOfBounds(board *[][]byte, x, y int) bool {
	var height, widht = len(*board), len((*board)[0])

	return x >= widht || x < 0 || y >= height || y < 0
}
