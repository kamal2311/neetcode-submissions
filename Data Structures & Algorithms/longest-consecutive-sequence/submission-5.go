func longestConsecutive(nums []int) int {

    if len(nums) == 0 {
        return 0
    }

    all := make(map[int]bool)

    for _, n := range nums {
        if !all[n] {
            all[n] = true
        }
    }   

    starts := []int{}

    for k, v := range nums {
        if !all[v-1] {
            starts = append( starts, k)
        }
    }

    maxLen := 1

    for _, i := range starts {
        len := 1
        next := nums[i] + 1
        for all[next] {
            next++
            len++
            if len > maxLen {
                maxLen = len
            }
        }        
    }

    return maxLen
}
