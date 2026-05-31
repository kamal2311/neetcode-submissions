func maxArea(heights []int) int {

	maxArea := 0
	i, j := 0, len(heights)-1
	for i < j {

		// Since the formula only depends on the minumum of the two ends, moving the other one will only always reduce the area, so we try the other choice. moving the bar that was the min of the two
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
