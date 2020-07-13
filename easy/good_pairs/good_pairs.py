class Solution(object):
    def numIdenticalPairs(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        total_pairs = 0
        values_processed = {}

        for i in nums:
            if not i in values_processed:
                occurrences = nums.count(i)
                if occurrences > 1:
                    all_pairs = occurrences * (occurrences - 1)
                    increasing_order_pairs = all_pairs / 2
                    total_pairs = total_pairs + increasing_order_pairs
                    values_processed[i] = True

        return int(total_pairs)


s = Solution()
print("Solution 1 : ", s.numIdenticalPairs([1, 2, 3, 1, 1, 3]))
print("Solution 2 : ", s.numIdenticalPairs([1, 1, 1, 1]))
print("Solution 3 : ", s.numIdenticalPairs([1, 2, 3]))
