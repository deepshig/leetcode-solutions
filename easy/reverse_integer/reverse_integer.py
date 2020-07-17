import numpy


class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        INTMAX32 = 2147483647
        if abs(x) > INTMAX32:
            return 0

        negative = False
        if x < 0:
            negative = True
            x = abs(x)

        str_x = str(x)
        reverse_str = str_x[::-1]
        reverse_int = int(reverse_str)

        if reverse_int > INTMAX32:
            return 0
        if negative:
            reverse_int = reverse_int * -1

        return reverse_int


s = Solution()
print("Solution 1 : ", s.reverse(123))
print("Solution 2 : ", s.reverse(-123))
print("Solution 3 : ", s.reverse(120))
