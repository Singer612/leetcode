package main

func isAnagram(s string, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}

func main() {

}
