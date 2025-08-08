package main

func combinationSum3(k int, n int) [][]int {
	index := 9
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(sum, start int)
	dfs = func(sum, start int) {
		if len(path) == k {
			if sum == n {
				tmp := make([]int, 0)
				tmp = append(tmp, path...)
				res = append(res, tmp)
			}
			return
		}
		for i := start; i <= index-(k-len(path))+1 && sum+i <= n; i++ {
			path = append(path, i)
			dfs(sum+i, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 1)
	return res
}
func main() {

}
