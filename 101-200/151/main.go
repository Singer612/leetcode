package main

import (
	"strings"
)

// 解法一
func reverseWords1(s string) string {
	words := make([]string, 0)
	for _, word := range strings.Split(s, " ") {
		if word != "" {
			words = append(words, word)
		}
	}

	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}

// 解法二
func reverseWords(s string) string {
	tmp := []byte(s)
	j := 0
	for i := 0; i < len(tmp); i++ {
		if tmp[i] != ' ' {
			// 除了字符串第一位，其他只要碰到非空格（也就是遇到单词了），先给j赋值空格
			if j != 0 {
				tmp[j] = ' '
				j++
			}
			// 将整个单词赋值给新数组
			for i < len(tmp) && tmp[i] != ' ' {
				tmp[j] = tmp[i]
				j++
				i++
			}
		}
	}
	// 有效的字符串
	tmp = tmp[:j]
	// 翻转整个字符串
	reverseString(tmp)
	// 翻转每个单词
	k := 0
	for i := 0; i <= len(tmp); i++ {
		if i == len(tmp) || tmp[i] == ' ' {
			reverseString(tmp[k:i])
			k = i + 1
		}
	}
	return string(tmp)
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	s := "a good   example"
	reverseWords(s)
}
