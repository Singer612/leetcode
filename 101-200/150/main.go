package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	for _, v := range tokens {
		temp, err := strconv.Atoi(v)
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
			stack = append(stack, temp)
		}
	}
	return stack[len(stack)-1]
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	result := evalRPN(tokens)
	fmt.Println(result)
}
