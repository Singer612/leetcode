package main

func main() {

}
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	res := 1
	prev := nums[1] - nums[0]
	if prev != 0 {
		res = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prev <= 0 || diff < 0 && prev >= 0 {
			res++
			prev = diff
		}
	}
	return res
}
