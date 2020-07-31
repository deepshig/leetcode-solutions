class Solution(object):
    def canPlaceFlowers(self, flowerbed, n):
        """
        :type flowerbed: List[int]
        :type n: int
        :rtype: bool
        """
        total_plots = len(flowerbed)
        for i in range(0, total_plots):
            left, right = False, False
            if i == 0 or flowerbed[i-1] == 0:
                left = True
            if i == total_plots-1 or flowerbed[i+1] == 0:
                right = True
            if flowerbed[i] == 0 and left and right:
                flowerbed[i] = 1
                n -= 1
            if n <= 0:
                return True
        return False


s = Solution()
print("Solution 1 : ", s.canPlaceFlowers([1, 0, 0, 0, 1], 1))
print("Solution 1 : ", s.canPlaceFlowers([1, 0, 0, 0, 1], 2))
