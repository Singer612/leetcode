package main

// 构建next数组
func getNext(next []int, s string) {
	j := 0
	next[0] = 0
	for i := 1; i < len(s); i++ {
		// 不相等
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		// 相等
		if s[i] == s[j] {
			j++
		}
		// 给next数组赋值
		next[i] = j
	}

}
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	j := 0
	next := make([]int, len(needle))
	getNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}

func main() {

}
