func trap(height []int) int {

	total := 0

	for i := 0; i < len(height); i++ {
		
		l := findLeft(height, i)
		r := findRight(height, i)

		water := min(height[l],height[r]) - height[i]
		total += water
	}
	
	return total

}

func findLeft(h []int, i int) int {

	maxIndex,max := i, h[i]

	for j := 0; j < i; j++ {
		if h[j] > max {			
			maxIndex = j
			max = h[j]
		}
	} 
	return maxIndex
}

func findRight(h []int, i int) int {
	
	maxIndex,max := i, h[i]

	for j := len(h) - 1; j > i; j-- {
		if h[j] > max {			
			maxIndex = j
			max = h[j]
		}
	}
	return maxIndex
}