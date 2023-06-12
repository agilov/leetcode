package word_search_ii

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
