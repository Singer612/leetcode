package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			return []int{v, i}
		}
		m[target-nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	twoSum(nums, 6)

}
