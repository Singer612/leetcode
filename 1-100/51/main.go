package main

import (
	"math"
)

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	path := make([]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == n {
			tmp := generateBoard(path, n)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if isValid(path, start, i) {
				path = append(path, i)
				dfs(start + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}
func isValid(path []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if col == path[i] || math.Abs(float64(row-i)) == math.Abs(float64(col-path[i])) {
			return false
		}
	}
	return true
}
func generateBoard(queens []int, n int) []string {
	board := make([]string, 0)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

func main() {

}
