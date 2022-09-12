package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	charCount := make(map[rune]int)

	for _, c := range magazine {
		charCount[c]++
	}

	for _, noteChar := range ransomNote {
		if left, ok := charCount[noteChar]; !ok || left == 0 {
			return false
		}

		charCount[noteChar]--
	}

	return true
}
