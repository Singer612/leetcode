package main

func isHappy(n int) bool {
	m := make(map[int]bool)
	//如果n不为1并且数值没有出现过，需要继续循环
	for n != 1 && !m[n] {
		m[n] = true
		n = getSum(n)
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		// 计算当前位数平方和
		sum = sum + (n%10)*(n%10)
		// 计算下一位
		n = n / 10
	}
	return sum
}

func main() {

}
