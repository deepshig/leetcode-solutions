class Solution(object):
    def shuffle(self, nums, n):
        """
        :type nums: List[int]
        :type n: int
        :rtype: List[int]
        """

        shuffled_array = []
        for i in range(0, n):
            shuffled_array.append(nums[i])
            shuffled_array.append(nums[i+n])

        return shuffled_array


s = Solution()
print("Solution 1 :", s.shuffle([2, 5, 1, 3, 4, 7], 3))
print("Solution 2 :", s.shuffle([1, 2, 3, 4, 4, 3, 2, 1], 4))
print("Solution 1 :", s.shuffle([1, 1, 2, 2], 2))
