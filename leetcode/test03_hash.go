package main

import (
	"fmt"
	"slices"
)

func longestConsecutive(nums []int) int {
	slices.Sort(nums)
	fmt.Println(nums)
	bucket := 0
	flag := 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			bucket = 1
		} else {
			if nums[i-1]+1 == nums[i] {
				flag++
			} else if nums[i-1] != nums[i] {
				flag = 1
			}
			if flag > bucket {
				bucket = flag
			}
		}
	}
	return bucket
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
