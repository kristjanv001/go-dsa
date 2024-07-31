// Given an array of integers, a start index, and an end index, reverse the integers in the
// array in-place without using any extra memor
//
// Leetcode 344

package array

func ReverseInPlace(list []int, start, end int) {
	for start < end {
		tmp := list[start]
		list[start] = list[end]
		list[end] = tmp

		start++
		end--
	}
}
