class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        max_sum = -float("inf")
        current_sum = 0
        for n in nums:
            current_sum = max(n, current_sum+n)
            max_sum = max(max_sum, current_sum)

        return max_sum


s = Solution()
print("Solution 1 :", s.maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4]))
