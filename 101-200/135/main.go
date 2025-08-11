package main

func main() {
	ratings := []int{1, 0, 2}
	candy(ratings)
}

func candy(ratings []int) int {
	candyArr := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			candyArr[i] = candyArr[i-1] + 1
		} else {
			candyArr[i] = 1
		}
	}
	for j := len(ratings) - 1; j >= 0; j-- {
		if j < len(ratings)-1 && ratings[j] > ratings[j+1] {
			candyArr[j] = max(candyArr[j+1]+1, candyArr[j])
		}
	}
	res := 0
	for _, v := range candyArr {
		res = res + v
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
