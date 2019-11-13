package main

import "fmt"

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	// 长度不一样，返回False
	sList := []rune(s)
	tList := []rune(t)
	if len(sList) != len(tList) {
		return false
	}
	// 1.统计所有字母出现次数
	sNums := countLetters(sList)
	tNums := countLetters(tList)
	// 2.对比两个集合元素是否一致
	for k, sv := range sNums { // 遍历sNums
		if tv, ok := tNums[k]; ok { // 判断sv是否在tNums中，值是否相等
			if sv != tv {
				return false
			}
		} else {
			return false
		}

	}
	return true
}

func countLetters(str []rune) map[rune]int {
	nums := make(map[rune]int)
	for _, s := range str {
		nums[s] = nums[s] + 1
	}
	return nums
}

// @lc code=end

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
