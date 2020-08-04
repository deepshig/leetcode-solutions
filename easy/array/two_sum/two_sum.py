from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        numberToIndex = {}
        for i, n in enumerate(nums):
            partner = target - n
            if partner in numberToIndex:
                return [numberToIndex[partner], i]
            numberToIndex[n] = i
        return


s = Solution()
print("Solution 1 : ", s.twoSum([2, 7, 11, 15], 9))
print("Solution 2 : ", s.twoSum([3, 3], 6))
