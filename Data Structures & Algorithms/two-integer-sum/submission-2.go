func twoSum(nums []int, target int) []int {
    
    valToInt := make(map[int]int)

    for i,v := range nums {
        if k, ok := valToInt[target-v]; ok {
            return []int{k, i}
        }

        valToInt[v] = i 
    }

    return []int{}
}
