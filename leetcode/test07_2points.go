package main

import (
	"fmt"
)

func trapx(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l, r := 0, len(height)-1
	lmax, rmax := 0, 0
	result := 0
	for l < r {
		if height[l] < height[r] {
			if height[l] > lmax {
				lmax = height[l]
			} else {
				result += lmax - height[l]
			}
			l++
		} else {
			if height[r] > rmax {
				rmax = height[r]
			} else {
				result += rmax - height[r]
			}
			r--
		}
	}
	return result
}

func trap(height []int) int {
	result := 0
	length := len(height)
	i := 0

	for i < length {
		// 如果当前高度为0，直接跳过
		if height[i] == 0 {
			i++
			continue
		}

		pkg := 0
		l := height[i]
		r := 0
		j := i + 1

		// 向右找第一个 >= l 且距离 >1 的柱子
		for ; j < length; j++ {
			if height[j] >= l && j-i > 1 {
				r = height[j]
				// 计算 [i, j] 区间的水量
				water := l*(j-i-1) - pkg
				if water > 0 {
					result += water
				}
				l = r
				i = j
				break
			}

			// 相邻且右柱 >= 左柱，不积水，直接移动 i
			if height[j] >= l && j-i == 1 {
				i++
				pkg = 0
				break
			}

			// 否则累加中间柱子高度（用于后续水量计算）
			pkg += height[j]
		}

		// 如果没找到右边界（r == 0），说明右侧没有比 l 高的柱子
		if r == 0 && j == length {
			// 找右侧最大值作为新的左边界
			tmp := 0
			for k := i + 1; k < length; k++ {
				if height[k] > tmp {
					tmp = height[k]
				}
			}

			if tmp == 0 {
				break
			}

			// 用右侧最大值替换当前高度（模拟左边界变低）
			height[i] = tmp

			// 避免死循环：如果和下一个相等，直接跳过
			if i+1 < length && height[i] == height[i+1] {
				i++
			}
		}
	}

	return result
}

func main() {
	//nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 1, 2, 1, 2, 1}
	//nums := []int{4, 2, 3}
	//nums := []int{5, 4, 1, 2}
	nums := []int{0, 1, 2, 0, 3, 0, 1, 2, 0, 0, 4, 2, 1, 2, 5, 0, 1, 2, 0, 2}
	//nums := []int{8, 5, 4, 1, 8, 9, 3, 0, 0}
	fmt.Println(nums)
	fmt.Println(trapx(nums))
}
