package main

import "sort"

func permuteUnique(nums []int) [][]int {
	res, path, used := make([][]int, 0), make([]int, 0), make([]bool, len(nums))
	sort.Ints(nums)
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for i := 0; i < len(nums); i++ {
			if i != 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				dfs(start + 1)
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	dfs(0)
	return res
}

func main() {
	nums := []int{1, 1, 2}
	permuteUnique(nums)
}
