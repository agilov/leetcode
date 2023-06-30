package ransom_note

import "testing"

func TestRansomNote(t *testing.T) {
	if canConstruct("aa", "ab") {
		t.Error("Should not construct aa from ab")
	}

	if !canConstruct("ab", "abc") {
		t.Error("Should construct ab from abc")
	}
}
