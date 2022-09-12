package roman_to_integer

var translationMap = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var result, operand int

	for _, char := range s {
		current := translationMap[char]
		if operand < current {
			operand = current - operand
		} else {
			result += operand
			operand = current
		}
	}

	result += operand
	return result
}
