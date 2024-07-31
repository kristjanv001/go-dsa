// Given a list of integers (1,2,...,n), find a duplicate number in O(n) time.

// O(n) time
// O(n) space
//
// create a hashmap
// check if num exists, if it does, we have found the duplicate
// if not, add it to hashmap
// if duplicate is not detected, return -1

package array

func FindDuplicate(list []int) int {
	hashmap := make(map[int]bool)

	for i := 0; i < len(list); i++ {
		if _, ok := hashmap[list[i]]; ok {
			return list[i]
		} else {
			hashmap[list[i]] = true
		}
	}

	return -1
}
