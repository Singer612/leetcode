package main

func isValid(s string) bool {
	stack := make([]int32, 0)
	m := map[int32]int32{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			if m[v] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	s := "([])"
	isValid(s)
}
