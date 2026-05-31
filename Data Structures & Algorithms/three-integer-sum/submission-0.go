func threeSum(nums []int) [][]int {

	res := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		twoSum(nums, i+1, -nums[i], &res)
	}

	return res

}

func twoSum(nums []int, start int, target int, res *[][]int) {

	for i, j := start, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == target {
			*res = append(*res, []int{-target, nums[i], nums[j]})
			i++
			j--
			for i < j && nums[i] == nums[i-1] {
				i++
			}
			for i < j && nums[j] == nums[j+1] {
				j--
			}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
}



