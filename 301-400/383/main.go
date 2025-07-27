package main

func canConstruct(ransomNote string, magazine string) bool {
	record := make([]int, 26)
	// 通过record数据记录 magazine里各个字符出现次数
	for _, v := range magazine {
		record[v-'a']++
	}
	// 遍历ransomNote，在record里对应的字符个数做--操作
	for _, v := range ransomNote {
		record[v-'a']--
		// 如果小于零说明ransomNote里出现的字符，magazine没有
		if record[v-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {

}
