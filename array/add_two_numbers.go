package array


// Given two numbers as an array like [2,9] and [9,9,9] return the sum of the numbers they
// represent like [1,0,2,8], because 29+999=1028.

// 1. convert both num slices to numbers
// 2. add those numbers to get the sum
// 3. break the sum to single digits (reverse process)

func AddTwoNumbers(num1, num2 []int) []int {
	intOne := convertDigitsToNum(num1)
	intTwo := convertDigitsToNum(num2)

	sum := intOne + intTwo

	return convertNumToDigits(sum)

}

func convertDigitsToNum(nums []int) int {
	var num int
	multiplier := 1

	for i := len(nums) - 1; i >= 0; i-- {
		num += nums[i] * multiplier
		multiplier *= 10
	}

	return num
}

func convertNumToDigits(num int) []int {
	digits := []int{}

	if num == 0 {
		return digits
	}

	for num > 0 {
		digit := num % 10
		digits = append([]int{digit}, digits...)
		num /= 10
	}

	return digits
}
