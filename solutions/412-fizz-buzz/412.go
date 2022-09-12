package fizz_buzz

import "strconv"

func fizzBuzz(n int) []string {
	const fizz = 3
	const buzz = 5

	fizzIndex := fizz
	buzzIndex := buzz

	result := make([]string, n)
	for i := 1; i <= n; i++ {
		str := ""
		if i == fizzIndex {
			str += "Fizz"
			fizzIndex += fizz
		}

		if i == buzzIndex {
			str += "Buzz"
			buzzIndex += buzz
		}

		if str == "" {
			str = strconv.Itoa(i)
		}

		result[i-1] = str
	}

	return result
}
