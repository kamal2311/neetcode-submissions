func trap(height []int) int {

	total := 0

	leftMaxes := make([]int, len(height))
	rightMaxes := make([]int, len(height))

	leftmax := -1
	for i := 0; i < len(height); i++ {		
		if height[i] > leftmax {
			leftmax = height[i]
		}
		leftMaxes[i] = leftmax
	}

	rightmax := -1
	for i := len(height) - 1; i >= 0 ; i-- {		
		if height[i] > rightmax {
			rightmax = height[i]
		}
		rightMaxes[i] = rightmax
	}

	for i := 0; i < len(height); i++ {
		water := min(leftMaxes[i],rightMaxes[i]) - height[i]
		total += water
	}
	
	return total

}

