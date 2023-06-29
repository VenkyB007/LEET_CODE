package leet

func SearchMatrix(matrix [][]int, target int) bool {
	targetRow := 0
	for i := 0; i < len(matrix); i++ {
		if target >= matrix[i][0] {
			targetRow = i
		}
	}

	for i := 0; i < len(matrix[0]); i++ {
		if target == matrix[targetRow][i] {
			return true
		}
	}
	return false
}
