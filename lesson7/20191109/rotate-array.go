package main

import "fmt"

func rotate(nums []int, k int) {
	numsLen := len(nums)
	kMod := k % numsLen
	doubleNums := append(nums, nums...)
	nums = doubleNums[numsLen-kMod : numsLen*2-kMod]
	fmt.Println("nums:", nums)
}

func main() {
	// https://leetcode-cn.com/problems/rotate-array/
	var nums []int
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
}
