func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for k, v := range nums {
        if i, ok := seen[target-v]; ok {
            return []int{i, k}
        }
        seen[v] = k
    }

    return []int{}
}
