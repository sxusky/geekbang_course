# -*- config:utf-8 -*-

class Solution(object):
    def rotate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        nums_len = len(nums)
        k_mod = k % nums_len
        for i in range(nums_len):
            nums.append(nums[i])
        nums = nums[nums_len-k_mod:nums_len-k_mod+nums_len]
        print(nums)

nums_a = [1,2,3,4,5,6,7]
s = Solution()
s.rotate(nums_a, 3)