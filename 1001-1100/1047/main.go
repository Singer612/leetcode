package main

func removeDuplicates(s string) string {
	stack := make([]int32, 0)
	for _, v := range s {
		if len(stack) != 0 && v == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}

func main() {

}
