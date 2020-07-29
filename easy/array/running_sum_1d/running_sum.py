class Solution(object):
    def runningSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        sums = []
        sums.append(nums[0])

        for i in range(1, len(nums)):
            sum = sums[i-1] + nums[i]
            sums.append(sum)

        return sums


s = Solution()
print("Solution 1 :", s.runningSum([1, 2, 3, 4]))
print("Solution 2 :", s.runningSum([1, 1, 1, 1, 1]))
print("Solution 3 :", s.runningSum([3, 1, 2, 10, 1]))
