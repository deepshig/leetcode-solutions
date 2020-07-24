class Solution(object):
    def canWinNim(self, n):
        """
        :type n: int
        :rtype: bool
        """
        return n % 4 != 0


s = Solution()
print("Solution 1 : ", s.canWinNim(4))
print("Solution 2 : ", s.canWinNim(5))
print("Solution 3 : ", s.canWinNim(8))
print("Solution 4 : ", s.canWinNim(16))
