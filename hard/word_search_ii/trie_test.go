package word_search_ii

import (
	"testing"
)

// go test -v -run=TestTrieImpl ./hard/word_search_ii/trie*.go
func TestTrieImpl(t *testing.T) {
	trie := newTrie()

	trie.add("foo")
	trie.add("bar")

	if !trie.find("fo") {
		t.Error("Cannot find foo")
	}

	if !trie.find("foo") {
		t.Error("Cannot find foo")
	}

	if !trie.find("bar") {
		t.Error("Cannot find bar")
	}

	if trie.find("baz") {
		t.Error("Found baz should not")
	}
}
