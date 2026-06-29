func findMin(nums []int) int {

    // we are trying to locate the pivot (the min)
    // if nums[m] > nums[r]
    // we know that pivot happened strictly to the right where numbers continued to increase and dropped to min at some point in that half
    // l = m + 1
    // else if nums[m] <= nums[r]
    // we can safely discard the right half but pivot could also be at m so
    // r = m

    l, r := 0, len(nums) - 1 
    for l < r {
        m := l + (r - l)/2
        if nums[m] > nums[r] {
            l = m + 1 
        } else {
            r = m
        }        
    }

    return nums[l]

}
