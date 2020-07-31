class Solution(object):
    def findPairs(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        occurrences = {}
        for n in nums:
            if n in occurrences:
                occurrences[n] += 1
            else:
                occurrences[n] = 1

        pairs = 0
        for n, o in occurrences.items():
            if k == 0 and o > 1:
                pairs += 1
            elif k > 0:
                partner = n + k
                if partner in occurrences and occurrences[partner] > 0:
                    pairs += 1

        return pairs


s = Solution()
print("Solution 1 :", s.findPairs([3, 1, 4, 1, 5], 2))
print("Solution 2 :", s.findPairs([1, 2, 3, 4, 5], 1))
print("Solution 3 :", s.findPairs([1, 3, 1, 5, 4], 0))
