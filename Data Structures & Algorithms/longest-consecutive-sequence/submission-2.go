func longestConsecutive(nums []int) int {

    all := make(map[int]bool)

    for _, n := range nums {       
        all[n] = true       
    }   

    maxLen := 0

    for _, num := range nums {
        streak, currNum := 0, num
        for all[currNum] {
            currNum++
            streak++
            if streak > maxLen {
                maxLen = streak
            }
        }        
    }

    return maxLen
}
