package words_data_structure

import (
	"agilov/leetcode/helpers"
	"testing"
)

type WordDictionaryTestCase struct {
	Operations   []string `json:"operations"`
	Arguments    []string `json:"arguments"`
	Expectations []bool   `json:"expectations"`
}

// go test -v -run=TestWordDictionary ./medium/words_data_structure/trie_recursive*.go ./medium/words_data_structure/dictionary.go
func TestWordDictionary(t *testing.T) {

	data, err := helpers.TestDataProvider[WordDictionaryTestCase]("./testdata/word_dictionary_test_data.json")

	if err != nil {
		t.Fatal(err)
	}

	var dict WordDictionary

	for i := range data {
		for j := range data[i].Operations {
			testWordDictionary(t, &data[i], &dict, j)
		}
	}
}

func testWordDictionary(t *testing.T, tc *WordDictionaryTestCase, dict *WordDictionary, j int) {
	operation := tc.Operations[j]
	word := tc.Arguments[j]
	expect := tc.Expectations[j]

	switch operation {
	case "WordDictionary":
		*dict = Constructor(NewTrieRecursive())
	case "addWord":
		dict.AddWord(word)
	case "search":
		result := dict.Search(word)
		if result != expect {
			t.Errorf("dict.Search(`%v`): expected - `%v` but got - `%v`", word, expect, result)
		}
	}
}

// go test -v -run=^$ -bench=BenchmarkWordDictionaryRecursiveTrie ./medium/words_data_structure/trie_recursive*.go ./medium/words_data_structure/dictionary.go
func BenchmarkWordDictionaryRecursiveTrie(b *testing.B) {
	data, err := helpers.TestDataProvider[WordDictionaryTestCase]("./testdata/word_dictionary_benchmark_data.json")
	if err != nil {
		b.Fatal(err)
	}
	tc := data[0]
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dict := Constructor(NewTrieRecursive())
		for j, operation := range tc.Operations {
			switch operation {
			case "addWord":
				dict.AddWord(tc.Arguments[j])
			case "search":
				dict.Search(tc.Arguments[j])
			}
		}
	}
}
