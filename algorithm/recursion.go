package algorithm

import "fmt"

//一个数组只能从左右两边拿，两个人依次拿，求最后谁拿到总和最大
func GetMaxNum(nums []int)  {
	a := maxNumProcessing1(nums, 0, len(nums) - 1)
	fmt.Printf("a:%d", a)
	b := maxNumProcessing2(nums, 0, len(nums) - 1)
	fmt.Printf("b:%d", b)
}

func maxNumProcessing1(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	return max(maxNumProcessing2(nums, left+1, right)+nums[left],maxNumProcessing2(nums, left , right-1) + nums[right])
}

func maxNumProcessing2(nums []int, left int, right int) int {
	if left == right {
		return 0
	}
	return min(maxNumProcessing1(nums, left+1, right),maxNumProcessing1(nums, left , right-1))
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min(num1 int, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}