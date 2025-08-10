package main

func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; {
		tmp := 0
		prev := prices[i]
		for i < len(prices)-1 && prices[i] < prices[i+1] {
			i++
			tmp = prices[i] - prev
		}
		res = res + tmp
		i++
	}
	return res
}

func main() {
	prices := []int{1, 2, 3, 4, 5}
	maxProfit(prices)
}
