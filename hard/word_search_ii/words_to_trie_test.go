package word_search_ii

import (
	"agilov/leetcode/helpers"
	"testing"
)

// go test -v -run=TestFindWordsWordsToTrie ./hard/word_search_ii/*.go
func TestFindWordsWordsToTrie(t *testing.T) {
	data, err := helpers.TestDataProvider[struct {
		Board  []string `json:"board"`
		Words  []string `json:"words"`
		Expect []string `json:"expect"`
	}]("./testdata/word_search_ii_testdata.json")

	if err != nil {
		t.Fatal(err)
	}

	for _, d := range data {
		result := findWords(convertBoard(d.Board), d.Words)

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

// go test -v -run=^$ -bench=BenchmarkFindWordsWordsToTrie ./hard/word_search_ii/*.go
func BenchmarkFindWordsWordsToTrie(b *testing.B) {
	data, _ := helpers.TestDataProvider[struct {
		Board  []string `json:"board"`
		Words  []string `json:"words"`
		Expect []string `json:"expect"`
	}]("./testdata/word_search_ii_bench.json")
	d := data[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findWords(convertBoard(d.Board), d.Words)
	}
}
