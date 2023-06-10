package implement_trie_prefix_tree

type TrieLoop struct {
	links [26]*Trie
	isEnd bool
}

func ConstructorTrieLoop() TrieLoop {
	return TrieLoop{}
}

func (this *TrieLoop) Insert(word string) {
	return
}

func (this *TrieLoop) Search(word string) bool {
	return false
}

func (this *TrieLoop) StartsWith(prefix string) bool {
	return false
}
