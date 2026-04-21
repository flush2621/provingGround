package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	length := len(s)
	ans, l := 0, 0
	for i := 0; i < length; i++ {
		if i == 0 {
			ans = 1
		}
		for j := l; j < i; j++ {
			if s[i] == s[j] {
				l = j + 1
				break
			}
		}
		if i-l+1 > ans {
			ans = i - l + 1
		}
	}
	return ans
}

func main() {
	s := "abcabcbb"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))
}
