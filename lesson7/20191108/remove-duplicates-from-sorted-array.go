package main

import "fmt"

func removeDuplicates(nums []int) int {
	// 双游标疗法
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
	var nums []int
	nums = []int{1, 1, 2}
	fmt.Println(nums, removeDuplicates(nums), nums)
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(nums, removeDuplicates(nums), nums)
}
