package sentence_similarity

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	pairs := make(map[string]bool, len(similarPairs)*2)

	if len(sentence1) != len(sentence2) {
		return false
	}

	for _, p := range similarPairs {
		pairs[p[0]+"_"+p[1]] = true
		pairs[p[1]+"_"+p[0]] = true
	}

	for i := 0; i < len(sentence1); i++ {
		if sentence1[i] == sentence2[i] {
			continue
		}

		if !pairs[sentence1[i]+"_"+sentence2[i]] && !pairs[sentence2[i]+"_"+sentence1[i]] {
			return false
		}
	}

	return true
}
