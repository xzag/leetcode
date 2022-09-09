package ransom_note

/**
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true


Constraints:

1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.
*/

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
