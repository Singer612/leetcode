package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	res, path := make([][]int, 0), make([]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		if start == len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			if i != start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

func main() {

}
