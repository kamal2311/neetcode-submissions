func hasDuplicate(nums []int) bool {
    
    seen := make(map[int]struct{})
    
    for _, v := range nums {
        if _, ok := seen[v]; ok {
            return true
        }
        seen[v] = struct{}{}
    }

    return false
}

// [1,2,4,5,2] -> true
// [1,2,3,4,5] -> false