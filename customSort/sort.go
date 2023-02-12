package customSort

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
// 每次循环都将最小值左移到第二层循环开始的地方
// 时间复杂度 最优情况 N^2, 最差情况 N^2 即 O(N^2)
// 空间复杂度 最优情况 1, 最差情况 1
func SelectionSort(nums []int) {
	lens := len(nums)
	for i := 0; i < lens; i++ {
		for j := i; j < lens; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// InsertionSort InsertSort 插入排序
// 第一层循环代表指针，第二循环负责将指针左侧保持有序
// 时间复杂度 最优情况 N, 最差情况 N^2 即 O(N^2)
// 空间复杂度 最优情况 1, 最差情况 1
func InsertionSort(nums []int) {
	lens := len(nums)
	for i := 1; i < lens; i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}

// MergeSort 归并排序
// 递归合并左右两个数组
// 时间复杂度 最优情况 N * LogN, 最差情况 N * LogN 即 O(N * LogN)
// 空间复杂度 最优情况 N, 最差情况 N
func MergeSort(nums []int) {
	lens := len(nums)
	if lens < 2 {
		return
	}
	mergeProcess(nums, 0, lens-1)
}

func mergeProcess(nums []int, l int, r int) {
	if l == r {
		return
	}
	mid := l + ((r - l) >> 1)
	mergeProcess(nums, l, mid)
	mergeProcess(nums, mid+1, r)
	merge(nums, mid, l, r)
}

func merge(nums []int, mid int, l int, r int) {
	var temp []int
	i := l       //左指针
	j := mid + 1 // 右指针
	for i <= mid && j <= r {
		if nums[i] < nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}
	for j <= r {
		temp = append(temp, nums[j])
		j++
	}
	for i <= mid {
		temp = append(temp, nums[i])
		i++
	}
	for t := 0; t < len(temp); t++ {
		nums[l+t] = temp[t]
	}
}

// QuickSort 快速排序
// 任意选择一个位置的数（这里选最后一个），以这个数为准将小于这个数的放左边，大于这个数的放右边，等于这个数的放中间，依次递归
// 时间复杂度 最优情况 LogN * LogN, 最差情况 N^2 即 O(N * LogN) 与target取值有关
// 空间复杂度 最优情况 logN, 最差情况 N 均值logN
func QuickSort(nums []int) {
	lens := len(nums)
	if lens < 2 {
		return
	}
	quickProcess(nums, 0, lens-1)
}
func quickProcess(nums []int, l int, r int) {
	if l >= r {
		return
	}
	i := l - 1 //左慢指针 指定小于部分
	f := l     //左快指针 指定等于部分
	j := r     //右指针 指定大于部分
	target := nums[r]
	for j >= f {
		if nums[f] > target {
			nums[f], nums[j] = nums[j], nums[f]
			j--
		}
		if nums[f] < target {
			i++
			nums[i], nums[f] = nums[f], nums[i]
			f++
		}
		if nums[f] == target {
			f++
		}
	}
	quickProcess(nums, l, i)
	quickProcess(nums, j+1, r)
}

//堆排序
