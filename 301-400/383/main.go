package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	m := make(map[int32]int)
	for _, v := range magazine {
		m[v]++
	}
	for _, v := range ransomNote {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}
	return true
}

func main() {

}
