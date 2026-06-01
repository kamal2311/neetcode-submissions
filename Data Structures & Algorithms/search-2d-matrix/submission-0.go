func searchMatrix(matrix [][]int, target int) bool {

	// first locate the row, 
	//  start with m/2 -> 
		//check if target > curr[0] and < curr[n-1]
		// 	you have found the row
		// else if target > curr[n-1]
		// search following half of the rows
		// if target < curr[0]
		// search the preceding half of the rows

	// then locate the col using standard 1D binary search within the row


	rows := len(matrix)
	cols := len(matrix[0])
	rowL, rowR := 0, rows - 1
	

	for rowL <= rowR {
		
		currR := rowL + (rowR - rowL) / 2

		if matrix[currR][0] <= target && target <= matrix[currR][cols-1] {
			return binSearch(matrix[currR], target)
		} else if target < matrix[currR][0] {
			rowR = currR - 1
		} else {
			rowL = currR + 1
		}
	}

	return false

}

func binSearch(nums []int, target int) bool {
	l, r := 0, len(nums) - 1

	for l <= r {
		curr := l + (r -l)/2
		if nums[curr] == target {
			return true
		} else if target < nums[curr] {
			r = curr - 1
		} else {
			l = curr + 1
		}
	}

	return false

}
