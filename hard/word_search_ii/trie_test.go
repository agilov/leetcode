package word_search_ii

import (
	"agilov/leetcode/helpers"
	"testing"
)

// go test -v -run=TestWordSearchII ./hard/word_search_ii/trie*.go

func TestWordSearchII(t *testing.T) {
	data, err := helpers.TestDataProvider[struct {
		Board  []string `json:"board"`
		Words  []string `json:"words"`
		Expect []string `json:"expect"`
	}]("./testdata/word_search_ii_testdata.json")

	if err != nil {
		t.Fatal(err)
	}

	for _, d := range data {
		result := findWords(d.Board, d.Words)

		if len(d.Expect) != len(result) {
			t.Fatalf("For board %v and words %v expected %v but got %v", d.Board, d.Words, d.Expect, result)
		}

		for _, word := range d.Expect {
			found := false
			for _, w := range result {
				if word == w {
					found = true
					break
				}
			}

			if !found {
				t.Fatalf("For board %v and words %v expected %v but got %v", d.Board, d.Words, d.Expect, result)
			}
		}
	}
}

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
