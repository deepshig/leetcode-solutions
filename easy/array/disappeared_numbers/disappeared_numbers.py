class Solution(object):
    def findDisappearedNumbers(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        i = 0
        while i < len(nums):
            n = nums[i]
            if nums[n-1] != n:
                nums[i], nums[n-1] = nums[n-1], n
            else:
                i += 1

        disappeared_numbers = []
        for i in range(0, len(nums)):
            if nums[i] != i+1:
                disappeared_numbers.append(i+1)

        return disappeared_numbers


s = Solution()
print("Solution 1 : ", s.findDisappearedNumbers([4, 3, 2, 7, 8, 2, 3, 1]))
