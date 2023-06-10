package implement_trie_prefix_tree

import "testing"

// go test -v ./medium/implement_trie_prefix_tree/recursive*.go

type TestCaseRecursive struct {
	insert string
	arg    string
	expect bool
}

var recursiveSearchDataProvider = [...]TestCaseRecursive{
	{"wort", "wor", false},
	{"wort", "wort", true},
	{"sword", "wor", false},
	{"leet", "le", false},
	{"leet", "leet", true},
}

// go test -v -run TestRecursiveInsertSearch ./medium/implement_trie_prefix_tree/recursive*.go
func TestRecursiveInsertSearch(t *testing.T) {
	for _, tc := range recursiveSearchDataProvider {
		trie := Constructor()
		trie.Insert(tc.insert)
		result := trie.Search(tc.arg)

		if result != tc.expect {
			t.Errorf("Search('%v') expected `%v`, but got `%v`", tc.arg, tc.expect, result)
		}
	}
}

// go test -v -run TestInsertPrefix ./medium/implement_trie_prefix_tree/recursive*.go
func TestInsertPrefix(t *testing.T) {
	cases := [...]TestCaseRecursive{
		{"wort", "wor", true},
		{"sword", "wor", false},
		{"leet", "le", true},
	}

	for _, tc := range cases {
		trie := Constructor()
		trie.Insert(tc.insert)
		result := trie.StartsWith(tc.arg)

		if result != tc.expect {
			t.Errorf("StartsWith('%v') expected `%v`, but got `%v`", tc.arg, tc.expect, result)
		}
	}
}

// go test -v -run=^$ -bench=BenchmarkTrieRecursive ./medium/implement_trie_prefix_tree/recursive*.go
func BenchmarkTrieRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trie := Constructor()

		for _, tc := range recursiveSearchDataProvider {
			trie.Insert(tc.insert)
			trie.Search(tc.arg)
			trie.StartsWith(tc.arg)
		}
	}
}
