package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	chars := make(map[rune]int)
	var current, max, start int
	for index, c := range s {
		if position, ok := chars[c]; ok && position >= start {
			start = position + 1
			current = index - position
		} else {
			current++
		}

		chars[c] = index

		if current > max {
			max = current
		}
	}

	return max
}
