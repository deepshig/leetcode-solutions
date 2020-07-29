class Solution(object):
    def maxProduct(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        first_max = max(nums)
        nums.remove(first_max)
        second_max = max(nums)
        return (first_max-1)*(second_max-1)


s = Solution()
print("Solution 1 : ", s.maxProduct([3, 4, 5, 2]))
print("Solution 2 : ", s.maxProduct([1, 5, 4, 5]))
print("Solution 3 : ", s.maxProduct([3, 7]))
print("Solution 4 : ", s.maxProduct([10, 2, 5, 2]))
