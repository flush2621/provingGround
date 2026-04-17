package main

import "fmt"

//func maxArea(height []int) int {
//	area := 0
//	mid := 0
//	for i := 0; i < len(height); i++ {
//		for j := i + 1; j < len(height); j++ {
//			if height[i] < height[j] {
//				mid = height[i] * (j - i)
//			} else {
//				mid = height[j] * (j - i)
//			}
//			if mid > area {
//				area = mid
//			}
//		}
//	}
//	return area
//}

func maxArea(height []int) int {
	area := 0
	left, right := 0, len(height)-1
	for left < right {
		width := right - left
		h := min(height[left], height[right])
		current := width * h
		if current > area {
			area = current
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}
