from typing import List


class Solution:
    def findErrorNums(self, nums: List[int]) -> List[int]:
        duplicate, missing = 0, 0
        for n in nums:
            n = abs(n)
            if nums[n-1] < 0:
                duplicate = n
            else:
                nums[n-1] *= -1

        for i in range(0, len(nums)):
            if nums[i] > 0:
                missing = i+1

        return [duplicate, missing]


s = Solution()
print("Solution 1 : ", s.findErrorNums([1, 2, 2, 4]))
