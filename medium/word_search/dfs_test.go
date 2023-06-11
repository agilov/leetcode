package word_search

import (
	"agilov/leetcode/helpers"
	"testing"
)

func convertBoard(board []string) [][]byte {
	result := make([][]byte, len(board))

	for i, word := range board {
		result[i] = make([]byte, len(word))
		for j, c := range word {
			result[i][j] = byte(c)
		}
	}

	return result
}

// go test -v -run=TestExistDfs ./medium/word_search/dfs*.go
func TestExistDfs(t *testing.T) {
	data, err := helpers.TestDataProvider[struct {
		Board  []string `json:"board"`
		Words  []string `json:"words"`
		Expect []bool   `json:"expect"`
	}]("./testdata/word_search_testdata.json")

	if err != nil {
		t.Fatal(err)
	}

	for _, d := range data {
		for i, word := range d.Words {
			result := exist(convertBoard(d.Board), word)

			if result != d.Expect[i] {
				t.Fatalf("For board %v and word %v expected %v but got %v", d.Board, word, d.Expect[i], result)
			}
		}
	}
}

// go test -v -run=^$ -bench=BenchmarkExistDfs ./medium/word_search/dfs*.go
func BenchmarkExistDfs(b *testing.B) {
	data, _ := helpers.TestDataProvider[struct {
		Board  []string `json:"board"`
		Words  []string `json:"words"`
		Expect []bool   `json:"expect"`
	}]("./testdata/word_search_benchdata.json")

	for i := 0; i < b.N; i++ {
		for _, word := range data[0].Words {
			exist(convertBoard(data[0].Board), word)
		}
	}
}
