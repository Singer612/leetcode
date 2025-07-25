package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := right + left/2
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] == target {
			return middle
		}
	}
	return -1
}

func main() {

}
