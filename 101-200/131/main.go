package main

func partition(s string) [][]string {
	res, path := make([][]string, 0), make([]string, 0)
	var dfs func(start int)
	dfs = func(start int) {
		// 如果起始位置等于s的大小，说明已经找到了一组分割方案了
		if start == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(s); i++ {
			str := s[start : i+1]
			if compare(str) {
				path = append(path, str)
				// 寻找i+1为起始位置的子串
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}
func compare(s string) bool {
	left, right := 0, len(s)-1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
func main() {

}
