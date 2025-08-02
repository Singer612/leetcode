package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := max(nums)
	left := nums[:idx]
	right := nums[idx+1:]
	root := &TreeNode{Val: nums[idx]}
	root.Left = constructMaximumBinaryTree(left)
	root.Right = constructMaximumBinaryTree(right)
	return root
}
func max(nums []int) int {
	res := 0
	maxNum := math.MinInt
	for i, v := range nums {
		if v > maxNum {
			maxNum = v
			res = i
		}
	}
	return res
}

func main() {

}
