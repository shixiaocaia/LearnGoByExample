package main

import "fmt"

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
	fmt.Println("冒泡排序结果：", nums)
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
	fmt.Println("选择排序结果：", nums)
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
	fmt.Println("插入排序结果：", nums)
}

func main() {
	nums := []int{4, 2, 3, 5, 6, 10}

	bubbleSort(nums)
	selectSort(nums)
	insertSort(nums)
}
