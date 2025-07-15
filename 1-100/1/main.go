package main

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]
*/

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] != target {
				continue
			} else {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, v := range nums {
		anotherNum := target - v
		if index, ok := hashMap[anotherNum]; ok {
			return []int{index, i}
		} else {
			hashMap[v] = i
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	// 创建一个哈希表，用来存放数组的元素值以及索引值
	// key 是数组中的元素值，value 是数组中元素值的索引
	hashMap := make(map[int]int)
	// 遍历整个数组
	for i, num := range nums {
		// 计算与当前元素配对的目标值
		anotherNum := target - num
		// 判断哈希表中是否存在配对的元素
		index, exists := hashMap[anotherNum]
		if exists {
			// 如果存在，返回两个数的索引
			return []int{index, i}
		} else {
			// 如果不存在，将当前元素及其索引添加到哈希表中
			hashMap[num] = i
		}
	}
	// 如果找不到配对的元素，返回空数组
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	//twoSum1(nums, target)
	twoSum(nums, target)
}
