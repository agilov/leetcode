package main

import "testing"

// go test -v ./implement_trie_prefix_tree_recursive*.go

type TestCase struct{
	insert string
	arg string
	expect bool
}

// go test -v -run TestInsertSearch ./implement_trie_prefix_tree_recursive*.go
func TestInsertSearch(t *testing.T) {
	cases := [...] TestCase {
		{"wort", "wor", false},
		{"wort", "wort", true},
		{"sword", "wor", false},
		{"leet", "le", false},
		{"leet", "leet", true},
	}

	for _, tc := range cases {
		trie := Constructor()
		trie.Insert(tc.insert)
		result := trie.Search(tc.arg)

		if result != tc.expect {
			t.Errorf("Search('%v') expected `%v`, but got `%v`", tc.arg, tc.expect, result)
		}
	}
}

// go test -v -run TestInsertPrefix ./implement_trie_prefix_tree_recursive*.go
func TestInsertPrefix(t *testing.T) {
	cases := [...] TestCase {
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