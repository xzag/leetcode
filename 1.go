package leetcode

import "sort"

/**
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.


Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
*/

func twoSum(nums []int, target int) []int {
	type composite struct {
		value int
		index int
	}

	data := make([]composite, len(nums), len(nums))
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
