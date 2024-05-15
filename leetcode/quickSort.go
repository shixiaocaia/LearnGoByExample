package main

import (
	"fmt"
	"math/rand"
)

func sortArray(nums []int) []int {
	var quickSort func(vec []int, left, right int)
	quickSort = func(vec []int, left, right int) {
		if left >= right {
			return
		}

		l, r := left, right
		piv := left + rand.Intn(right-left+1)
		pivNum := vec[piv]
		vec[right], vec[piv] = vec[piv], vec[right]
		for i := left; i < right; i++ {
			if vec[i] < pivNum {
				vec[i], vec[left] = vec[left], vec[i]
				left++
			}
		}
		vec[left], vec[right] = vec[right], vec[left]
		quickSort(vec, l, left-1)
		quickSort(vec, left+1, r)
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func main() {
	vec := []int{2, 3, 1, 4}
	fmt.Println(sortArray(vec))
}
