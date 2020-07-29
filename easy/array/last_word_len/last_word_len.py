class Solution(object):
    def lengthOfLastWord(self, s):
        """
        :type s: str
        :rtype: int
        """
        s = s.rstrip()
        i = len(s)-1
        length_of_word = 0
        while i >= 0 and s[i] != ' ':
            length_of_word += 1
            i -= 1
        return length_of_word


s = Solution()
print("Solution 1 : ", s.lengthOfLastWord("Hello World"))
