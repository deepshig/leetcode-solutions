class Solution(object):
    def toLowerCase(self, str):
        """
        :type str: str
        :rtype: str
        """
        for c in str:
            if (c >= 'A') and (c <= 'Z'):
                str = str.replace(c, chr(ord(c) + 32), 1)
        return str


s = Solution()
print("Solution 1 :", s.toLowerCase("Hello"))
print("Solution 2 :", s.toLowerCase("here"))
print("Solution 3 :", s.toLowerCase("LOVELY"))
