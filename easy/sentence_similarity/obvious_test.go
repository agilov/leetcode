package sentence_similarity

import "testing"

// go test -v -run=TestAreSentencesSimilar ./easy/sentence_similarity/*.go
func TestAreSentencesSimilar(t *testing.T) {
	s1 := []string{"great", "acting", "skills"}
	s2 := []string{"fine", "drama", "talent"}
	pairs := [][]string{
		[]string{"great", "fine"},
		[]string{"drama", "acting"},
		[]string{"skills", "talent"},
	}
	expect := true
	ans := areSentencesSimilar(s1, s2, pairs)

	if ans != expect {
		t.Errorf("For sentences %v and %v and pairs %v result should be %v but got %v", s1, s1, pairs, expect, ans)
	}
}
