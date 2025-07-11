package main

import "fmt"

func moveZeroes(nums []int) {
	i, j := 0, 0
	length := len(nums)
	for i < length {
		if length == 1 && nums[i] == 0 {
			return
		}
		if nums[i] == 0 {
			i++
		}
		if nums[i] != 0 {
			nums[j] = nums[i]
			i++
			j++
		}
	}
}

func main() {
	var nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}
