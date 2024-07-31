// Given an array of numbers like {1, 2, -4, 6, 3} returns unique triplets from the numbers
// with sum that equals zero like {-4, 1, 3}.
//
// Leetcode 15

package array

import "sort"

func ZeroSumTriplets(list []int) [][]int {
	result := [][]int{}

	if len(list) < 3 {
		return result
	}

	sort.Ints(list)

	for i := 0; i < len(list); i++ {
		if i > 0 && list[i] == list[i-1] {
			continue
		}

		left := i + 1
		right := len(list) - 1

		for left < right {
			current := list[left] + list[right] + list[i]

			if current == 0 {
				result = append(result, []int{list[i], list[left], list[right]})
				left++

				for list[left] == list[left-1] && left < right {
					left++
				}

			} else if current > 0 {
				right--
				continue

			} else {
				left++
				continue
			}
		}

	}

	return result
}
