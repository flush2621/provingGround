package main

import "fmt"

func moveZeroes(nums []int) {
	var bucket []int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			bucket = append(bucket, i)
		}
	}

	for j := 0; j < len(bucket); j++ {
		for k := bucket[j] - j; k < len(nums); k++ {
			if k == len(nums)-1 {
				nums[k] = 0
				break
			}
			nums[k] = nums[k+1]
		}
	}
	fmt.Println(nums)
}

func moveZeroesX(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
	fmt.Println(nums)
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroesX(nums)
}
