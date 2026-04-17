package main

import (
	"fmt"
	"slices"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	hashMap := make(map[string][]string)
	for _, str := range strs {
		sorted := sortString(str)
		hashMap[sorted] = append(hashMap[sorted], str)
	}
	for _, group := range hashMap {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	chars := []byte(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func groupAnagramsX(strs []string) (result [][]string) {
	hashMap := make(map[string][]string)
	for _, s := range strs {
		sToBytes := []byte(s)
		slices.Sort(sToBytes)
		hashMap[string(sToBytes)] = append(hashMap[string(sToBytes)], s)
	}
	for _, value := range hashMap {
		result = append(result, value)
	}
	return
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagramsX(strs)

	fmt.Println(result)
	// 输出: [[bat] [nat tan] [ate eat tea]] 或类似顺序
}
