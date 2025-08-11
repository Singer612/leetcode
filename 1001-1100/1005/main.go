package main

import (
	"sort"
)

func main() {

}
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			if k > 0 {
				nums[i] = -nums[i]
				k--
			}
		}
	}
	if k > 0 {
		if k%2 == 1 {
			sort.Ints(nums)
			nums[0] = -nums[0]
		}
	}
	for i := 0; i < len(nums); i++ {
		res = res + nums[i]
	}
	return res
}
