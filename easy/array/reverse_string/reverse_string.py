class Solution(object):
    def reverseString(self, s):
        """
        :type s: List[str]
        :rtype: None Do not return anything, modify s in-place instead.
        """
        i, j = 0, len(s)-1
        while i < j:
            s[i], s[j] = s[j], s[i]
            i += 1
            j -= 1
        return


s = Solution()
a = ["h", "e", "l", "l", "o"]
s.reverseString(a)
print("Solution 1 : ", a)

b = ["H", "a", "n", "n", "a", "h"]
s.reverseString(b)
print("Solution 2 : ", b)
