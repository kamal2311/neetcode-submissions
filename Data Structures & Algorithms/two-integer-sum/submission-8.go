func twoSum(nums []int, target int) []int {
    
    indexedNums := make([][2]int, len(nums))

    for i,v := range nums {
        indexedNums[i] = [2]int{v,i}
    }

    sort.Slice(indexedNums, func (i, j int) bool {
        return indexedNums[i][0] < indexedNums[j][0]
    })

    l,r := 0, len(indexedNums) - 1

    for l < r {
        sum := indexedNums[l][0] + indexedNums[r][0]
        if sum == target {

            if indexedNums[l][1] < indexedNums[r][1]{
                return []int{indexedNums[l][1], indexedNums[r][1]} 
            }            
            return []int{indexedNums[r][1], indexedNums[l][1]}            
        } else if sum < target {
            l++
        } else {
            r--
        }
    }

    return nil
}
