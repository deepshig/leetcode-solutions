class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        profit = 0
        for i in range(1, len(prices)):
            if prices[i] > prices[i-1]:
                profit += prices[i] - prices[i-1]
        return profit


s = Solution()
print("Solution 1 : ", s.maxProfit([7, 1, 5, 3, 6, 4]))
print("Solution 2 : ", s.maxProfit([1, 2, 3, 4, 5]))
print("Solution 3 : ", s.maxProfit([7, 6, 4, 3, 1]))
print("Solution 4 : ", s.maxProfit([1, 7, 2, 3, 6, 7, 6, 7]))
