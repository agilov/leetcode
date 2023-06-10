package main

type Trie struct {
	links [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	if word == "" {
		return
	}

	next := this.links[word[0]-'a']

	if next == nil {
		next = &Trie{}
		this.links[word[0]-'a'] = next
	}

	if len(word) == 1 {
		next.isEnd = true
		return
	}

	next.Insert(word[1:])
}

func (this *Trie) Search(word string) bool {
	next := this.links[word[0]-'a']

	if word == "" || next == nil {
		return false
	}

	if len(word) == 1 {
		return next.isEnd
	}

	return next.Search(word[1:])
}

func (this *Trie) StartsWith(prefix string) bool {
	next := this.links[prefix[0]-'a']

	if prefix == "" || next == nil {
		return false
	}

	return len(prefix) == 1 || next.StartsWith(prefix[1:])
}
