package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	dict := [26]int{}

	for i := 0; i < len(magazine); i++ {
		dict[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		idx := ransomNote[i] - 'a'
		if dict[idx] > 0 {
			dict[idx]--
		} else {
			return false
		}
	}

	return true
}
