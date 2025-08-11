package main

func main() {

}

func jump(nums []int) int {
	res, maxReach, end := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxReach = max(maxReach, i+nums[i])
		if i == end {
			end = maxReach
			res++
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
