// Given an array of integers A, construct a new array B such that B[i] = product of all items
// in A except A[i] without using division in O(n) time.
//
// This problem is based on the prefix sum pattern
//
// Leetcode 238

package array

func ProductOfAllOtherElements(list []int) []int {
	n := len(list)

	if n < 1 {
		return []int{}
	}

	prefix := make([]int, n)
	suffix := make([]int, n)
	result := []int{}

	// create prefix
	prefix[0] = list[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * list[i]
	}

	// create suffix
	suffix[n-1] = list[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * list[i]
	}

	// construct result array
	for i := 0; i < n; i++ {
		if i == 0 {
			result = append(result, suffix[i+1])
		} else if i == n-1 {
			result = append(result, prefix[i-1])
		} else {
			result = append(result, prefix[i-1]*suffix[i+1])
		}
	}

	return result
}
