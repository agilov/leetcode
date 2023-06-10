package implement_trie_prefix_tree

type TrieLoop struct {
	links [26]*TrieLoop
	isEnd bool
}

func ConstructorTrieLoop() TrieLoop {
	return TrieLoop{}
}

func (this *TrieLoop) Insert(word string) {
	current := this

	for _, c := range word {
		if current.links[c-'a'] == nil {
			current.links[c-'a'] = &TrieLoop{}
		}

		current = current.links[c-'a']
	}

	current.isEnd = true
}

func (this *TrieLoop) Search(word string) bool {
	current := this
	for _, c := range word {
		if current.links[c-'a'] == nil {
			return false
		}
		current = current.links[c-'a']
	}

	return current.isEnd
}

func (this *TrieLoop) StartsWith(prefix string) bool {
	current := this

	for _, c := range prefix {
		if current.links[c-'a'] == nil {
			return false
		}
		current = current.links[c-'a']
	}

	return true
}
