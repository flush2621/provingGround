package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0)

	// 用于去重的 map
	seen := make(map[string]bool)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} else if nums[i] > 0 {
			break
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				// 生成唯一 key
				key := fmt.Sprint(nums[i], nums[left], nums[right])

				// 检查是否已存在，不存在才添加
				if !seen[key] {
					seen[key] = true
					result = append(result, []int{nums[i], nums[left], nums[right]})
				}

				left++
				right--
			}
		}
	}

	return result
}

func threeSumX(nums []int) [][]int {
	slices.Sort(nums)
	length := len(nums)
	ans := make([][]int, 0)
	for i, num := range nums {
		if num > 0 {
			break
		}
		if i > 0 && nums[i-1] == num {
			continue
		}
		target := -num
		left, right := i+1, length-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				ans = append(ans, []int{num, nums[left], nums[right]})
				left++
				right--
			}
		}
	}
	return ans
}

func tx(nums []int) [][]int {
	slices.Sort(nums)
	length := len(nums)
	result := make([][]int, 0)
	for i, num := range nums {
		if num > 0 {
			break
		} else if i > 0 && num == nums[i-1] {
			continue
		}
		target := -num
		l, r := i+1, length-1
		for l < r {
			if target < nums[l]+nums[r] {
				r--
			} else if target > nums[l]+nums[r] {
				l++
			} else {
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				result = append(result, []int{num, nums[l], nums[r]})
				l++
				r--
			}
		}
	}
	return result
}

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{0, 0, 0, 0}
	//nums := []int{-100, -70, -60, 110, 120, 130, 160}
	nums := []int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}
	//nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	fmt.Println(tx(nums))
}
