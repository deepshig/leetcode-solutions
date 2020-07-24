class Solution(object):
    def isPerfectSquare(self, num):
        """
        :type num: int
        :rtype: bool
        """
        if num == 1:
            return True
        low, high = 1, num//2
        while low <= high:
            mid = (low+high)//2
            product = mid * mid
            if product == num:
                return True
            if product < num:
                low = mid + 1
            else:
                high = mid - 1
        return False


s = Solution()
print("Solution 1 : ", s.isPerfectSquare(16))
print("Solution 1 : ", s.isPerfectSquare(14))
print("Solution 1 : ", s.isPerfectSquare(1))
print("Solution 1 : ", s.isPerfectSquare(9))
print("Solution 1 : ", s.isPerfectSquare(25))
