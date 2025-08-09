package main

func findSubsequences(nums []int) [][]int {
	res, path := make([][]int, 0), make([]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) > 1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		if len(nums) == start {
			return
		}
		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				used[nums[i]] = true
				path = append(path, nums[i])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

func main() {
	nums := []int{4, 4, 3, 2, 1}
	findSubsequences(nums)
}
