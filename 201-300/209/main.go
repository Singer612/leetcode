package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	sum := 0                        // 记录滑动窗口的和
	length, res := 0, math.MaxInt64 // 滑动窗口的长度
	i := 0
	for j := 0; j < len(nums); j++ {
		sum = sum + nums[j]
		for sum >= target {
			// 记录下当前滑动窗口的长度
			length = j - i + 1
			res = min(length, res)
			sum = sum - nums[i]
			i++
		}
	}
	if res == math.MaxInt64 {
		return 0
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	minSubArrayLen(7, nums)
}
