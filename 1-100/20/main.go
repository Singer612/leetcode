package main

import "fmt"

func isValid(s string) bool {
	// 判断是否为奇数，必定不会有效
	if len(s)%2 == 1 {
		return false
	}
	// 构建一个左右括号的map，key为右括号，因为每次都是拿出右括号进行匹配
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 接下来构建数组，如果是左括号，就放入栈中，如果是右括号就拿出匹配
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		// map键值不存在，放入栈中
		if pairs[s[i]] <= 0 {
			stack = append(stack, s[i])
		} else {
			// 右括号，需要拿出匹配,如果栈里没有元素的话，也不有效
			if len(stack) != 0 && stack[len(stack)-1] == pairs[s[i]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}

}

func main() {
	s := "(("
	flag := isValid(s)
	fmt.Println(flag)
}
