package words_data_structure

type TrieLoop struct {
	head *TrieLoopNode
}

func NewTrieLoop() *TrieLoop {
	return &TrieLoop{head: &TrieLoopNode{}}
}

type TrieLoopNode struct {
	isEnd bool
	links [26]*TrieLoopNode
}

func (t *TrieLoop) add(word string) {

	current := t.head

	for _, c := range word {
		index := c - 'a'
		if current.links[index] == nil {
			current.links[index] = &TrieLoopNode{}
		}
		current = current.links[index]
	}

	current.isEnd = true
}

func (t *TrieLoop) search(word string) bool {
	return searchNode(word, t.head)
}

func searchNode(word string, node *TrieLoopNode) bool {
	if word == "" {
		return node.isEnd
	}
	current := node

	for i, c := range word {
		if c == '.' {
			for _, n := range current.links {
				if n != nil && searchNode(word[i+1:], n) {
					return true
				}
			}

			return false
		}

		index := c - 'a'
		if current.links[index] == nil {
			return false
		}
		current = current.links[index]
	}

	return current.isEnd
}
