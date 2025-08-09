package main

func permute(nums []int) [][]int {
	res, path, used := make([][]int, 0), make([]int, 0), make([]bool, len(nums))
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				dfs(start + 1)
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

func main() {

}
