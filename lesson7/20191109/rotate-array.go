package main

import "fmt"

func rotate(nums []int, k int) {
	// 拼接为两倍数组，k求mod，然后从len的位置往左挪kMod位，然后切片len个元素
	numsLen := len(nums)
	kMod := k % numsLen
	doubleNums := append(nums, nums...)
	nums = doubleNums[numsLen-kMod : numsLen-kMod+numsLen]
	fmt.Println("nums:", nums)
}

func main() {
	// https://leetcode-cn.com/problems/rotate-array/
	var nums []int
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
	rotate(nums, 0)
	rotate(nums, 1)
	rotate(nums, 4)
}
