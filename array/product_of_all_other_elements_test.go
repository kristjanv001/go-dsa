package array

import (
	"slices"
	"testing"
)

/*
TestProductOfAllOtherElements tests solution(s) with the following signature and problem description:

	ProductOfAllOtherElements(list []int) []int

Given an array of integers A, construct a new array B such that B[i] = product of all items
in A except A[i] without using division in O(n) time.
*/
func TestProductOfAllOtherElements(t *testing.T) {
	tests := []struct {
		list     []int
		products []int
	}{
		{[]int{}, []int{}},
		{[]int{2, 3}, []int{3, 2}},
		{[]int{1, 2, 3}, []int{6, 3, 2}},
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
		{[]int{0, 0, 1, 2}, []int{0, 0, 0, 0}},             // Multiple zeros in the array
		{[]int{1, 0, 3, 4}, []int{0, 12, 0, 0}},            // Single zero in the array
		{[]int{1, 2, 0, 4}, []int{0, 0, 8, 0}},             // Zero in the middle
		{[]int{0, 2, 3, 0}, []int{0, 0, 0, 0}},             // Zeros at both ends
		{[]int{0, 1, 2, 3, 4, 0}, []int{0, 0, 0, 0, 0, 0}}, // Leading and trailing zeros
	}

	for i, test := range tests {
		if got := ProductOfAllOtherElements(test.list); !slices.Equal(got, test.products) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.products, got)
		}
	}
}
