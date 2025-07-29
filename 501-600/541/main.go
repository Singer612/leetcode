package main

import "fmt"

func reverseStr(s string, k int) string {
	tmp := []byte(s)
	for i := 0; i < len(tmp); i += 2 * k {
		left, right := i, min(i+k-1, len(s)-1)
		reverseString(tmp, left, right)
	}
	return string(tmp)
}
func reverseString(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	s := "abcdefg"
	fmt.Println(reverseStr(s, 2))
}
