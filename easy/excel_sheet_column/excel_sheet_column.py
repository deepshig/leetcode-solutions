class Solution(object):
    def convertToTitle(self, n):
        """
        :type n: int
        :rtype: str
        """
        column_name = ""
        while n > 0:
            letter_number = n % 26
            if letter_number == 0:
                letter_number = 26
                n = n - 26
            letter = ord('A') + letter_number - 1
            column_name = chr(letter) + column_name
            n = int(n / 26)

        return column_name


s = Solution()
print("Solution 1 :", s.convertToTitle(1))
print("Solution 2 :", s.convertToTitle(28))
print("Solution 3 :", s.convertToTitle(701))
print("Solution 4 :", s.convertToTitle(52))
print("Solution 5 :", s.convertToTitle(703))
