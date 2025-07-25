package main

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	i, j := 0, len(nums)-1
	for k := len(nums) - 1; k >= 0; k-- {
		if nums[i]*nums[i] >= nums[j]*nums[j] {
			res[k] = nums[i] * nums[i]
			i++
		} else {
			res[k] = nums[j] * nums[j]
			j--
		}
	}
	return res
}

func main() {
	nums := []int{-7, -3, 2, 3, 11}
	sortedSquares(nums)
}
