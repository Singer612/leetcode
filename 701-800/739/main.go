package main

// 双循环超时
//func dailyTemperatures(temperatures []int) []int {
//	answer := make([]int, len(temperatures))
//	for i := 0; i < len(temperatures); i++ {
//		for j := i + 1; j < len(temperatures); j++ {
//			if temperatures[j] > temperatures[i] {
//				answer[i] = j - i
//				break
//			}
//		}
//	}
//	return answer
//}

func dailyTemperatures1(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return ans
}
func dailyTemperatures(temperatures []int) []int {
	var stack []int
	answer := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			answer[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return answer
}
func main() {

}
