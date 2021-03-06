package nowcoder

//给定一个二维数组，其每一行从左到右递增排序，从上到下也是递增排序。给定一个数，判断这个数是否在该二维数组中。
//
//Consider the following matrix:
//[
//[1,   4,  7, 11, 15],
//[2,   5,  8, 12, 19],
//[3,   6,  9, 16, 22],
//[10, 13, 14, 17, 24],
//[18, 21, 23, 26, 30]
//]
//
//Given target = 5, return true.
//Given target = 20, return false.

func isExist(matrix [][]int, i int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	c := cols - 1
	r := 0
	for c >= 0 && r <= rows-1 {
		if matrix[r][c] == i {
			return true
		}
		if matrix[r][c] < i {
			r++
		} else {
			c--
		}
	}
	return false
}
