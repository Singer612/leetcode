package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	// 设置两个set
	set1 := map[int]struct{}{}
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	set2 := map[int]struct{}{}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	// 遍历较短的那个set
	if len(set1) > len(set2) {
		// 交换顺序
		set1, set2 = set2, set1
	}
	for v := range set1 {
		if _, ok := set2[v]; ok {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	a, b := 3, 1
	a, b = b, a
	fmt.Println(a, b)
}
