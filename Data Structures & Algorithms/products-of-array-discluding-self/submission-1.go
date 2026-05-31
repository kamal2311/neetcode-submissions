func productExceptSelf(nums []int) []int {

    
    out := make([]int, len(nums))
    
    out[0] = 1
    out[len(nums)-1] = 1

    for i :=1; i < len(nums); i++ {
        out[i] = out[i-1] * nums[i-1] 
    }    

    right := 1 
    for j := len(nums) - 1; j >= 0; j-- {
        out[j] = out[j] * right
        right = right * nums[j]
    }

    return out

}
