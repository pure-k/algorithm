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

//堆概念：类似完全二叉树。大根堆：每个节点的子节点都小于等于该节点；小根堆：每个节点的子节点都大于等于该节点；
//特性：某节点的左子节点为2i+1；右节点为2i+2；父节点（i-1）/2 （向下取整）
//HeapSort堆排序
// 时间复杂度 O(N * LogN)
// 空间复杂度 O(1)
func HeapSort(nums []int)  {
	lens := len(nums)
	if lens < 2 {
		return
	}
	//堆化
	for i := 0; i < lens; i++{
		heapInsert(nums, i)
	}
	for heapSize := lens; heapSize > 0; {
		nums[0], nums[heapSize - 1] = nums[heapSize - 1], nums[0]
		heapSize --
		heapify(nums, 0, heapSize)
	}
}

func HeapSort2(nums []int)  {
	lens := len(nums)
	if lens < 2 {
		return
	}
}

//堆化，某个数在index位置，能否往下移动
func heapify(nums []int, index int, heapSize int)  {
	left := index * 2 + 1
	for left < heapSize {
		largest := index
		if nums[largest] < nums[left] {
			largest = left
		}
		if left + 1 < heapSize && nums[largest] < nums[left + 1] {
			largest = left + 1
		}
		if largest == index {
			break
		}
		nums[index], nums[largest] = nums[largest], nums[index]
		index = largest
		left = index * 2 + 1
	}
}
//某个数在index位置，能否往上移动
func heapInsert(nums []int, index int)  {
	for nums[index] > nums[(index-1)/2] {
		nums[index],nums[(index-1)/2] = nums[(index-1)/2],nums[index]
		index = (index-1)/2
	}
}

//计数排序(桶排序的特殊情况)
//适用于数据都在一个较小范围的数组
// 时间复杂度 O(N+M)  M指范围值
// 空间复杂度 O(N+M)
func CountingSort(nums []int)  {
	lens := len(nums)
	if lens < 2 {
		return
	}
	//找最小值，最大值
	max := nums[0]
	min := nums[0]
	for i := 1; i < lens; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	//创建计数数组
	var countLens int = max - min + 1
	countArr := make([]int, countLens)
	for j := 0; j < lens; j++ {
		countArr[nums[j] - min] ++
	}

	//数据还原
	index := 0
	for i := 0; i < countLens; i++ {
		for ;countArr[i] > 0; countArr[i] --{
			nums[index] = min + i
			index ++
		}
	}
}



//基数排序
//按进制分类，如十进制的分十个类，先将最右位分类，依次向左位分类
func RadixSort()  {
	
}

//桶排序
func BucketSort()  {
	
}
