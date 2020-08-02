from typing import List


class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        number_to_index = {}
        for i, n in enumerate(nums):
            if n in number_to_index and i-number_to_index[n] <= k:
                return True
            number_to_index[n] = i
        return False


s = Solution()
print("Solution 1 : ", s.containsNearbyDuplicate([1, 2, 3, 1], 3))
print("Solution 2 : ", s.containsNearbyDuplicate([1, 0, 1, 1], 1))
print("Solution 3 : ", s.containsNearbyDuplicate([1, 2, 3, 1, 2, 3], 2))
