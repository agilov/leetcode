package words_data_structure

type TrieRecursive struct {
	head *TrieRecursiveNode
}

func NewTrieRecursive() *TrieRecursive {
	return &TrieRecursive{head: &TrieRecursiveNode{}}
}

func (t *TrieRecursive) add(word string) {
	t.head.add(word)
}

func (t *TrieRecursive) search(word string) bool {
	return t.head.search(word)
}

type TrieRecursiveNode struct {
	isEnd bool
	links [26]*TrieRecursiveNode
}

func (tn *TrieRecursiveNode) add(word string) {
	if word == "" {
		tn.isEnd = true
		return
	}

	index := word[0] - 'a'

	if tn.links[index] == nil {
		tn.links[index] = &TrieRecursiveNode{}
	}

	tn.links[index].add(word[1:])
}

func (tn *TrieRecursiveNode) search(word string) bool {
	if word == "" {
		return tn.isEnd
	}

	if word[0] == '.' {
		for _, ch := range tn.links {
			if ch != nil && ch.search(word[1:]) {
				return true
			}
		}

		return false
	}

	character := tn.links[word[0]-'a']

	return character != nil && character.search(word[1:])
}
