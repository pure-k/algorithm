package customSort

import (
	"math/rand"
	"sort"
	"testing"
)

//冒泡排序准确性校验
//go test -v -run TestBubbleSort
func TestBubbleSort(t *testing.T) {
	testSort(BubbleSort, t)
}

func TestSelectionSort(t *testing.T) {
	testSort(SelectionSort, t)
}

func TestInsertionSort(t *testing.T) {
	testSort(InsertionSort, t)
}

func TestMergeSort(t *testing.T) {
	testSort(MergeSort, t)
}

func TestQuickSort(t *testing.T) {
	testSort(QuickSort, t)
}
func TestHeapSort(t *testing.T) {
	testSort(HeapSort, t)
}

func TestCountingSort(t *testing.T) {
	testSort(CountingSort, t)
}

//冒泡排序性能测试
//go test -bench BenchmarkBubbleSort
func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := generateRandomArray()
		BubbleSort(nums)
	}
}

//选择排序性能测试
func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := generateRandomArray()
		SelectionSort(nums)
	}
}

//插入排序性能测试
func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := generateRandomArray()
		InsertionSort(nums)
	}
}

//归并排序性能测试
func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := generateRandomArray()
		MergeSort(nums)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := generateRandomArray()
		QuickSort(nums)
	}
}
func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := generateRandomArray()
		HeapSort(nums)
	}
}

func BenchmarkCountingSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := generateRandomArray()
		CountingSort(nums)
	}
}
//系统排序性能测试
//go test -bench BenchmarkSort
func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := generateRandomArray()
		sort.Ints(nums)
	}
}

//随机数组
func generateRandomArray() []int {
	lens := rand.Intn(10000)
	var ret []int
	for i := 0; i < lens; i++ {
		ret = append(ret, rand.Intn(10000))
	}
	return ret
}

//验证其他排序方法准确性
func testSort(sortFunc func(nums []int), t *testing.T) {
	for i := 0; i < 100; i++ {
		nums := generateRandomArray()
		nums1 := make([]int, len(nums), len(nums))
		copy(nums1, nums)
		sortFunc(nums)
		sort.Ints(nums1)
		if !compareArray(nums, nums1, t) {
			t.Error("比对失败")
		}
	}
}

//比较两个数组
func compareArray(nums1 []int, nums2 []int, t *testing.T) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			t.Errorf("nums1[%d]值为%d,nums2[%d]值为%d", i, nums1[i], i, nums2[i])
			return false
		}
	}
	return true
}
