func search(nums []int, target int) int {

	l,r := 0, len(nums) - 1

	for l <= r {
		curr := l + (r-l)/2
		if nums[curr] == target {
			return curr
		} else if nums[curr] < target {
			l = curr + 1
		} else {
			r = curr - 1
		}
	}

	return -1

}
