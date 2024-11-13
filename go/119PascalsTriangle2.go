package main

//				1
//	           11
//			   121
//			  1331
//		  	 14641
func pascalsTriangle2(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}
	result := make([][]int, rowIndex+1)
	result[0] = []int{1}
	result[1] = []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		currentLength := i + 1
		currentRow := make([]int, currentLength)
		currentRow[0] = 1
		currentRow[currentLength-1] = 1
		for j := 0; j < currentLength-2; j++ {
			currentRow[j+1] = result[i-1][j] + result[i-1][j+1]
		}
		result[i] = currentRow
	}
	return result[rowIndex]
}
