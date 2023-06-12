package word_search_ii

type TrieNode struct {
	isEnd      bool
	word       string
	liknsCount uint8
	links      [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func (tn *TrieNode) has(char byte) bool {
	return tn.get(char) != nil
}

func (tn *TrieNode) get(char byte) *TrieNode {
	return tn.links[char-'a']
}

func (tn *TrieNode) getIsEnd(char byte) bool {
	return tn.links[char-'a'].isEnd
}

func (t *Trie) add(word string) {
	current := t.root

	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if current.links[idx] == nil {
			current.links[idx] = newTrieNode()
			current.liknsCount++
		}
		current = current.links[idx]
	}

	current.isEnd = true
	current.word = word
}

func (t *Trie) find(word string) bool {
	current := t.root

	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if current.links[idx] == nil {
			return false
		}
		current = current.links[idx]
	}

	return true
}

func newTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

func newTrieNode() *TrieNode {
	return &TrieNode{links: [26]*TrieNode{}, isEnd: false, liknsCount: 0}
}
