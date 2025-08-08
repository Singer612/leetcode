package main

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res, path := make([][]int, 0), make([]int, 0)
	var dfs func(start, sum int)
	dfs = func(start, sum int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		if len(path) == len(candidates) {
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i]+sum > target {
				break
			}
			if i != start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			dfs(i+1, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return res
}

func main() {

}
