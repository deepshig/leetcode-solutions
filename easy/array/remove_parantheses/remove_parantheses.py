class Solution(object):
    def removeOuterParentheses(self, S):
        """
        :type S: str
        :rtype: str
        """
        incomplete_paranthese = 0
        start = 1
        final_string = ""

        for i in range(0, len(S)):
            if S[i] == '(':
                incomplete_paranthese += 1
            else:
                incomplete_paranthese -= 1
            if incomplete_paranthese == 0:
                end = i
                final_string = final_string + S[start:end]
                start = i + 2

        return final_string


s = Solution()
print("Solution 1 :", s.removeOuterParentheses("(()())(())"))
print("Solution 2 :", s.removeOuterParentheses("(()())(())(()(()))"))
print("Solution 3 :", s.removeOuterParentheses("()()"))
