package words_data_structure

type Trie interface {
	add(word string)
	search(word string) bool
}

type WordDictionary struct {
	trie Trie
}

func Constructor(trie Trie) WordDictionary {
	return WordDictionary{trie: trie}
}

func (wd *WordDictionary) AddWord(word string) {
	wd.trie.add(word)
}

func (wd *WordDictionary) Search(word string) bool {
	return wd.trie.search(word)
}
