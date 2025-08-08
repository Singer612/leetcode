package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	path, res := make([]int, 0), make([][]int, 0)
	sort.Ints(candidates)
	var dfs func(sum, start int)
	dfs = func(sum, start int) {
		if sum == target {
			tmp := make([]int, 0)
			tmp = append(tmp, path...)
			res = append(res, tmp)
			return
		}
		if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			dfs(sum+candidates[i], i)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return res
}

func main() {

}
