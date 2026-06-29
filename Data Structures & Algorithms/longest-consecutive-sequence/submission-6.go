func longestConsecutive(nums []int) int {
    // store the numbers in a hashmap
    // initial answer = 0
    // iterate over the array one by one
    // for each index i, keep checking if nums[i] + 1 is in the hashmap, if yes, keep incrementing the counter.
    // that will be the len of a sequece starting at i, update the longest answer
    // How do I avoid repeating the calculation?
    // ignore the numbers that cannot be a start ( the ones which don't have n-1 in the map)

    all := make(map[int]struct{})

    for _, n := range nums {
        all[n] = struct{}{}
    }

    res := 0

    for _,n := range nums {
        if _, ok := all[n-1]; !ok{
            len := calcLen(n, all)
            if len > res {
                res = len
            }
        }
    }

    return res

}

func calcLen(n int, m map[int]struct{}) int{
    len := 1
    curr := n + 1    
    
    for  {
        if _, ok := m[curr]; ok {
            len++
            curr++
        } else {
            break
        }
    }
    return len
}
