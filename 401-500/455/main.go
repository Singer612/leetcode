package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res, j := 0, 0
	for i := 0; i < len(g); i++ {
		for j < len(s) {
			if s[j] >= g[i] {
				res++
				j++
				break
			} else {
				j++
			}
		}
	}
	return res
}

func main() {

}
