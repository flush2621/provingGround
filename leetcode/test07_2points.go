package main

import (
	"fmt"
)

func trapx(height []int) int {
	result := 0
	length := len(height)
	for i := 0; i < length; {
		pkg := 0
		l, r := height[i], 0
		j := i + 1
		//fmt.Println(l, r, i, j)
		for ; j < length; j++ {
			//fmt.Println(i, j, l)
			if height[j] >= l && l > 0 && j-i > 1 {
				r = height[j]
				//fmt.Println(pkg)
				if l*(j-i-1)-pkg > 0 {
					//fmt.Println(i, j, l*(j-i-1)-pkg)
					result += l*(j-i-1) - pkg
				}
				l = r
				i = j
				break
			}
			pkg += height[j]
			if height[j] >= l && l > 0 && j-i == 1 {
				l = height[j]
				pkg = 0
				i++
			}
		}
		//fmt.Println("f:", i, r)
		if r == 0 {
			if l > 0 {
				tmp := 0
				for k := i + 1; k < length; k++ {
					if height[k] > tmp {
						tmp = height[k]
					}
				}
				//fmt.Println(tmp, i, l)
				if tmp > 0 {
					height[i] = tmp
				}

				if i >= length-2 || tmp == 0 {
					break
				}
				if height[i] == height[i+1] {
					i++
				}
			} else {
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
	//nums := []int{0, 1, 2, 0, 3, 0, 1, 2, 0, 0, 4, 2, 1, 2, 5, 0, 1, 2, 0, 2}
	nums := []int{8, 5, 4, 1, 8, 9, 3, 0, 0}
	fmt.Println(nums)
	fmt.Println(trapx(nums))
}
