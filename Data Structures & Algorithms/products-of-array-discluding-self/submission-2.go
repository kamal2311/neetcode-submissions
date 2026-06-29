func productExceptSelf(nums []int) []int {

	// prod[i] = preProd[i] * postProd[i]
	// preProd[i] = 1* ...*nums[i-1]

	// postProd[i] = 1 *...*nums[i+1]

	res := make([]int, len(nums))

	pre := calcPre(nums)
	post := calcPost(nums)


	for i := 0; i < len(res); i++{
		res[i] = pre[i] * post[i]
	}

	return res

}

func calcPre(n []int) []int{

	pre := make([]int,len(n) + 2) // 0,0,0,0
	pre[0] = 1  
	pre[len(n)] = 1
	

	for i := 0; i < len(n); i++ {
		pre[i+1] = n[i] * pre[i] 
		
	}
	return pre[:len(n)]

}

func calcPost(n []int) []int{
	post := make([]int,len(n) + 2)
	post[0] = 1
	post[len(n)] = 1

	for i := len(n)-1; i >= 0; i-- {
		post[i] = n[i] * post[i+1]
	}

	return post[1:len(n)+1]
}
