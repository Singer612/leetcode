package main

func main() {

}

func canJump(nums []int) bool {
	cover := 0
	n := len(nums) - 1
	for i := 0; i <= cover; i++ { // 每次与覆盖值比较
		cover = max(i+nums[i], cover) //每走一步都将 cover 更新为最大值
		if cover >= n {
			return true
		}
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
