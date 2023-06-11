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


func findWords(board []string, words []string) []string {
	trie := newTrie()

	trie.add("test")

	return []string{"abd"}
}

// Inserting board to trie from concrete letter.
// Each letter have at most 4 neighbours, if neigbour is not inserted before - it marged as inserted after insertion
// insertion stops when there is no not-inserted letters left
func insert(board []string, trie *Trie, x int, y int) {

}
