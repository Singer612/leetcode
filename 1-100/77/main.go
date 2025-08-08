package main

func combine(n int, k int) [][]int {
	path, res := make([]int, 0), make([][]int, 0)
	var dfs func(n, k, start int)
	dfs = func(n, k, start int) {
		if len(path) == k {
			tmp := make([]int, 0)
			tmp = append(tmp, path...)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			dfs(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(n, k, 1)
	return res
}

func main() {

}
