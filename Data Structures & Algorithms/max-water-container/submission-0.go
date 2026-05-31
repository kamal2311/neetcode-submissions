func maxArea(heights []int) int {

	maxArea := 0
	i, j := 0, len(heights)-1
	for i < j {
		area := (j-i) * min (heights[i], heights[j])
		if area > maxArea {
			maxArea = area
		}
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea

}
