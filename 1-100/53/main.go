package main

import "math"

func main() {

}
func maxSubArray(nums []int) int {
	res := 0
	maxVal := math.MinInt
	for i := 0; i < len(nums); i++ {
		res = res + nums[i]
		maxVal = max(maxVal, res)
		if res < 0 {
			res = 0
		}
	}
	return maxVal
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
