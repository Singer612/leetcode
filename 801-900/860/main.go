package main

func main() {
	bills := []int{5, 5, 5, 10, 20}
	lemonadeChange(bills)
}

func lemonadeChange(bills []int) bool {
	fiveCount, tenCount := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			fiveCount++
		} else if bill == 10 {
			if fiveCount > 0 {
				fiveCount--
				tenCount++
			} else {
				return false
			}
		} else {
			if fiveCount > 0 && tenCount > 0 {
				tenCount--
				fiveCount--
			} else if fiveCount >= 3 {
				fiveCount = fiveCount - 3
			} else {
				return false
			}
		}
	}
	return true
}
