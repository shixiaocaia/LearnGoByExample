package main

import (
	"fmt"
	"math/rand"
)

// bubbleSort 冒泡排序
func bubbleSort(nums []int) {
	// 冒泡排序， 比较交换，稳定算法，时间复杂度O(n^2)，空间复杂度O(1)
	// 每一轮将较大的数放到最后，较小的数向前移动
	// 形成后部有序区
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	//fmt.Println("冒泡排序结果：", nums)
}

// selectSort 选择排序
func selectSort(nums []int) {
	// 选择排序，比较交换，不稳定算法，时间复杂度O(n^2)，空间复杂度O(1)
	// 每一轮选择最小的数放到前面
	// 形成前部有序区
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	//fmt.Println("选择排序结果：", nums)
}

// insertSort 插入排序
func insertSort(nums []int) {
	// 插入排序，比较移动，稳定算法，时间复杂度O(n^2)，空间复杂度O(1)
	// 从后往前比较，比当前数大的向后移动
	// 形成前部有序区
	// 数据规模小或者基本有序时，效率高
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > tmp {
			// 比当前数大的向后移动
			nums[j+1] = nums[j]
			j--
		}
		// tmp大于当前数，插入到当前数后面
		nums[j+1] = tmp
	}
	//fmt.Println("插入排序结果：", nums)
}
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	povit := partition(nums, left, right)
	povitl := povit - 1
	for povitl > 0 && nums[povitl] == nums[povitl+1] {
		povitl--
	}
	quickSort(nums, left, povitl)
	povitr := povit + 1
	for povitr < len(nums) && nums[povitr] == nums[povitr-1] {
		povitr++
	}
	quickSort(nums, povitr, right)
}

func partition(nums []int, left, right int) int {
	pivotIndex := rand.Intn(right-left) + left
	pivot := nums[pivotIndex]
	// pivot 完整遍历下
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
	//fmt.Println(nums)
	// 调整pivot位置
	nums[left], nums[right] = nums[right], nums[left]
	//fmt.Println(nums, pivotIndex)
	return left
}

func main() {
	nums := []int{5, 1, 1, 2, 0, 0}
	fmt.Println("before sort: ", nums)
	//bubbleSort(nums)
	//selectSort(nums)
	//insertSort(nums)
	quickSort(nums, 0, len(nums)-1)

	fmt.Println("after sort: ", nums)
}
