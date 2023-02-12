package customSort

import "fmt"

// BubbleSort 冒泡排序
// 将大值一点点向右推
// 时间复杂度 最优情况 N^2, 最差情况 N^2 即 O(N^2)
// 空间复杂度 最优情况 1, 最差情况 1
func BubbleSort(nums []int) {
	lens := len(nums)
	for i := 0; i < lens; i++ {
		for j := 0; j < lens-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// SelectionSort 选择排序
func SelectionSort() {

}

// InsertionSort InsertSort 插入排序
func InsertionSort() {

}

// MergeSort 归并排序
func MergeSort() {
	fmt.Println("dss")
}

//快速排序
func QuickSort() {

}

//堆排序
