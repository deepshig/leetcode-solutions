import copy


class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        number_of_zeroes = 0
        for i in range(0, len(nums)):
            if nums[i] == 0:
                number_of_zeroes += 1
            else:
                nums[i-number_of_zeroes] = nums[i]
        while number_of_zeroes > 0:
            nums[len(nums)-number_of_zeroes] = 0
            number_of_zeroes -= 1
        return


s = Solution()

nums = [0, 1, 0, 3, 12]
s.moveZeroes(nums)
print("Solution 1 : ", nums)

nums = [0, 0, 1]
s.moveZeroes(nums)
print("Solution 2 : ", nums)

nums = [1]
s.moveZeroes(nums)
print("Solution 3 : ", nums)
