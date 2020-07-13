class Solution(object):
    def kidsWithCandies(self, candies, extraCandies):
        """
        :type candies: List[int]
        :type extraCandies: int
        :rtype: List[bool]
        """
        max_candies = max(candies)
        can_have_max_candies = []

        for i in range(0, len(candies)):
            candies_needed = max_candies - candies[i]
            if candies_needed <= extraCandies:
                can_have_max_candies.append(True)
            else:
                can_have_max_candies.append(False)

        return can_have_max_candies


s = Solution()
print("Solution 1 : ", s.kidsWithCandies([2, 3, 5, 1, 3], 3))
print("Solution 2 : ", s.kidsWithCandies([4, 2, 1, 1, 2], 1))
print("Solution 3 : ", s.kidsWithCandies([12, 1, 12], 10))
