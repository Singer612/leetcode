package main

func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	nums = nums[:j]
	return j
}

func main() {

}
