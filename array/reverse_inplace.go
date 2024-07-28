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
