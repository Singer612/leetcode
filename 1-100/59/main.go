package main

func generateMatrix(n int) [][]int {
	// 定义上下左右边界
	top, bottom, left, right := 0, n-1, 0, n-1
	// 当前计数
	num := 1
	//一共需要填充进数组的数
	tar := n * n
	// 定义二维数组
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	// 数量不够时，一直计算
	for num <= tar {
		//上边一行
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		// 右边一列
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		// 下边一行
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		// 左边一列
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}

func main() {

}
