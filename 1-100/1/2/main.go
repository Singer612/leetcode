package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			res := []int{i, m[nums[i]]}
			return res
		}
		m[target-nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	twoSum(nums, 6)

}
