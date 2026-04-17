package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				flag := []int{i, j}
				return flag
			}
		}
	}
	return nil
}

func twoSumX(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if j, exists := hashMap[complement]; exists {
			return []int{j, i}
		}
		hashMap[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	//result := twoSum(nums, target)
	result := twoSumX(nums, target)
	fmt.Println(result)
}
