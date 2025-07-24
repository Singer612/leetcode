package main

// 自己无解哈哈哈哈哈，还是别自己想了，净歪门邪道

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
}

func (h *hp) Pop() interface{} {
	tmp := h.IntSlice
	res := tmp[len(tmp)-1]
	h.IntSlice = tmp[:len(tmp)-1]
	return res

}
func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// 存储和对应的键值对
	sumMap := make(map[int][][]int)
	sum := &hp{sort.IntSlice{}}
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			sumMap[v1+v2] = append(sumMap[v1+v2], []int{v1, v2})
			heap.Push(sum, v1+v2)
		}
	}
	res := make([][]int, 0)
	num := 0
	a := sum.IntSlice
	for len(a) != 0 {
		tmp := heap.Pop(sum).(int)
		length := len(sumMap[tmp])
		if num+length <= k {
			res = append(res, sumMap[tmp]...)
			num = num + length
		} else {
			res = append(res, sumMap[tmp][:k-num]...)
		}
	}
	return res
}

func main() {

}
