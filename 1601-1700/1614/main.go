package main

import "fmt"

func maxDepth(s string) int {
	stack := []byte{}
	maxDep := 0
	need1, need2 := '(', ')'
	for i := 0; i < len(s); i++ {
		if s[i] == byte(need1) {
			stack = append(stack, s[i])
		}
		if s[i] == byte(need2) {
			if len(stack) > maxDep {
				maxDep = len(stack)
			}
			stack = stack[:len(stack)-1]

		}
	}
	return maxDep
}

func main() {
	s := "(1)+((2))+(((3)))"
	fmt.Println(maxDepth(s))
}
