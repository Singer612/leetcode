package main

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	canCompleteCircuit(gas, cost)
}
func canCompleteCircuit(gas []int, cost []int) int {
	curTank, totalTank, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		curChange := gas[i] - cost[i]
		curTank = curTank + curChange
		totalTank = totalTank + curChange
		if curTank < 0 {
			start = i + 1
			curTank = 0
		}
	}
	if totalTank >= 0 {
		return start
	}
	return -1
}
