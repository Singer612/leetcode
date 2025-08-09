package main

func solveSudoku(board [][]byte) {
	var dfs func() bool
	dfs = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}
				for k := '1'; k <= '9'; k++ {
					if isValid(i, j, byte(k), board) {
						board[i][j] = byte(k)
						// 递归调用 dfs，继续填充下一个空位。
						if dfs() {
							// 如果递归调用返回 true，说明找到了一个完整的解，
							// 此时所有递归都应该向上返回 true。
							return true
						}
						board[i][j] = '.'
					}
				}
				// 如果 '1' 到 '9' 所有数字都尝试过，但都没有找到解，
				// 说明当前路径无效，向上回溯，返回 false。
				return false
			}
		}
		return true
	}
	dfs()
}
func isValid(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		// 检查当前行
		if board[row][i] == k {
			return false
		}
		// 检查当前列
		if board[i][col] == k {
			return false
		}
	}

	// 检查 3x3 小方格是否已存在相同的数字 k
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
