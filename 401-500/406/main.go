package main

import "sort"

func main() {

}
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			// 从大到小
			return people[i][0] > people[j][0]
		}
		// 从小到大
		return people[i][1] < people[j][1]
	})

	res := make([][]int, 0)
	for _, p := range people {
		index := p[1]
		res = append(res[:index], append([][]int{p}, res[index:]...)...)
	}
	return res
}
