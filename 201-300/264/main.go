package main

import (
	"container/heap"
	"sort"
)

var factors = []int{2, 3, 5}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func nthUglyNumber(n int) int {
	// 先把丑数1加进去
	h := &hp{sort.IntSlice{1}}
	// 用来标记已经出现过的丑数
	seen := map[int]struct{}{
		1: {},
	}
	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		if i == n {
			return x
		}
		for _, f := range factors {
			next := x * f
			if _, has := seen[next]; !has {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}
}

func main() {

}
