package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res, path := make([]string, 0), make([]string, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == 4 {
			if start == len(s) {
				res = append(res, strings.Join(path, "."))
				return
			}
		}
		for i := start; i < len(s); i++ {
			str := s[start : i+1]
			if isValid(str) {
				path = append(path, str)
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}
func isValid(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	tmp, _ := strconv.Atoi(s)
	if tmp < 0 || tmp > 255 {
		return false
	}
	return true
}
func main() {
	s := "0000"
	restoreIpAddresses(s)
}
