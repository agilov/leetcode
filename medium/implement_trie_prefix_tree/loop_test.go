package implement_trie_prefix_tree

import "testing"

// go test -v ./medium/implement_trie_prefix_tree/loop*.go

type TestCaseTrieLoop struct {
	insert string
	arg    string
	expect bool
}

var testTrieLoopInsertSearchDataProvider = [...]TestCaseTrieLoop{
	{"wort", "wor", false},
	{"wort", "wort", true},
	{"sword", "wor", false},
	{"leet", "le", false},
	{"leetcode", "le", false},
	{"leet", "leet", true},
}

// go test -v -run TestTrieLoopInsertSearch ./medium/implement_trie_prefix_tree/loop*.go
func TestTrieLoopInsertSearch(t *testing.T) {
	for _, tc := range testTrieLoopInsertSearchDataProvider {
		trie := ConstructorTrieLoop()
		trie.Insert(tc.insert)
		result := trie.Search(tc.arg)

		if result != tc.expect {
			t.Errorf("Search('%v') expected `%v`, but got `%v`", tc.arg, tc.expect, result)
		}
	}
}

var testTrieLoopInsertPrefixDataProvider = [...]TestCaseTrieLoop{
	{"wort", "wor", true},
	{"sword", "wor", false},
	{"leetcode", "leetco", true},
	{"leet", "le", true},
}

// go test -v -run TestTrieLoopInsertPrefix ./medium/implement_trie_prefix_tree/loop*.go
func TestTrieLoopInsertPrefix(t *testing.T) {
	for _, tc := range testTrieLoopInsertPrefixDataProvider {
		trie := ConstructorTrieLoop()
		trie.Insert(tc.insert)
		result := trie.StartsWith(tc.arg)

		if result != tc.expect {
			t.Errorf("StartsWith('%v') expected `%v`, but got `%v`", tc.arg, tc.expect, result)
		}
	}
}

// go test -run=^$ -v -bench=BenchmarkTrieLoop ./medium/implement_trie_prefix_tree/loop*.go
func BenchmarkTrieLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trie := ConstructorTrieLoop()

		for _, tc := range testTrieLoopInsertSearchDataProvider {
			trie.Insert(tc.insert)
			trie.Search(tc.arg)
			trie.StartsWith(tc.arg)
		}
	}
}
