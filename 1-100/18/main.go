package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	// 对数组进行排序
	sort.Ints(nums)
	for j := 0; j < len(nums); j++ {
		// 去重
		if j > 0 && nums[j] == nums[j-1] {
			continue
		}
		for i := j + 1; i < len(nums); i++ {
			// 定义快慢指针
			left, right := i+1, len(nums)-1
			// 对a去重
			if i > j+1 && nums[i] == nums[i-1] {
				continue
			}
			for left < right {
				tmp := nums[j] + nums[i] + nums[left] + nums[right]
				// 双指针逻辑，根据值的大小移动指针
				if tmp > target {
					right--
				} else if tmp < target {
					left++
				} else {
					tmp1 := []int{nums[j], nums[i], nums[left], nums[right]}
					res = append(res, tmp1)
					// 如果找到一个符合的结果，就对b去重，时刻保证left < right
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					// 对c去重，同时时刻保证left < right
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return res
}

func main() {

}
