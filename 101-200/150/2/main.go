package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, v := range tokens {
		tmp, err := strconv.Atoi(v)
		if err != nil {
			y, x := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				stack = append(stack, x+y)
			case "-":
				stack = append(stack, x-y)
			case "*":
				stack = append(stack, x*y)
			case "/":
				stack = append(stack, x/y)
			}
		} else {
			stack = append(stack, tmp)
		}
	}
	return stack[len(stack)-1]
}

func main() {

}
