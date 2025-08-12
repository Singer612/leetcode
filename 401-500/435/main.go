package main

import "sort"

func main() {

}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	res := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[i-1][1] {
			continue
		} else {
			res++
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
		}
	}
	return res
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
