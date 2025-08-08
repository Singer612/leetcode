package main

import "strconv"

var strMap = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	var path []byte
	res := make([]string, 0)
	var dfs func(digits string, start int)
	dfs = func(digits string, start int) {
		// 如果路径长度和digits长度相同，说明已经遍历结束
		if len(path) == len(digits) {
			// 可以将path放进结果数组
			res = append(res, string(path))
			return
		}
		num, _ := strconv.Atoi(string(digits[start]))
		str := strMap[num]
		for j := 0; j < len(str); j++ {
			path = append(path, str[j])
			dfs(digits, start+1)
			path = path[:len(path)-1]
		}
	}
	dfs(digits, 0)
	return res
}

func main() {

}
