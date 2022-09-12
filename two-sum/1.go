package two_sum

import "sort"

func twoSum(nums []int, target int) []int {
	stored := make(map[int]int, len(nums))

	for i, x := range nums {
		needToSolve := target - x
		if storedIndex, ok := stored[needToSolve]; ok {
			return []int{storedIndex, i}
		}

		stored[x] = i
	}

	return nil
}

func twoSum_WithSort(nums []int, target int) []int {
	type composite struct {
		value int
		index int
	}

	data := make([]composite, len(nums))
	for i, x := range nums {
		data[i] = composite{
			value: x,
			index: i,
		}
	}

	sort.Slice(data, func(i int, j int) bool {
		return data[i].value < data[j].value
	})

	left := 0
	right := len(nums) - 1

	for left < right {
		sum := data[left].value + data[right].value
		if sum == target {
			result := []int{data[left].index, data[right].index}
			sort.Ints(result)
			return result
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return nil
}
