func minEatingSpeed(piles []int, h int) int {

	// check the min between 1 and max(piles) for k such that condition of being able to eat is satisfied

	lo, hi := 1, findMax(piles)

	res := hi

	for lo <= hi {
		curr := lo + (hi - lo) / 2
		currHours := hoursToEat(piles, curr)		
		if currHours  <= h {
			res = curr			
			hi = curr - 1
		} else {
			lo = curr + 1
		}
	}

	return res

}

func hoursToEat(piles []int, rate int) int {
	total := 0
	for _, p := range piles{
		total += int(math.Ceil(float64(p)/float64(rate)))
	}
	return total
}

func findMax(piles []int) int{
	max := -1
	for _, p := range piles{
		if p > max {
			max = p
		}
	}
	return max
}
