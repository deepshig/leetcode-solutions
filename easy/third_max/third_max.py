class Solution(object):
    def thirdMax_first_solution(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        first_max = second_max = third_max = -float("inf")
        for val in nums:
            if val in (first_max, second_max, third_max):
                continue
            elif val > first_max:
                third_max, second_max, first_max = second_max, first_max, val
            elif val > second_max:
                third_max, second_max = second_max, val
            elif val > third_max:
                third_max = val

        if third_max > -float("inf"):
            return third_max

        return first_max

    def third_max_second_solution(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        first_max = max(nums)
        nums = list(filter(lambda x: x != first_max, nums))

        if len(nums) != 0:
            second_max = max(nums)
            nums = list(filter(lambda x: x != second_max, nums))

            if len(nums) != 0:
                third_max = max(nums)
                return third_max

        return first_max


s = Solution()
print("Solution 1 : ", s.thirdMax_first_solution([3, 2, 1]))
print("Solution 2 : ", s.thirdMax_first_solution([1, 2]))
print("Solution 3 : ", s.thirdMax_first_solution([2, 2, 3, 1]))
