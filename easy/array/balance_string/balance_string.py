class Solution(object):
    def balancedStringSplit(self, s):
        """
        :type s: str
        :rtype: int
        """
        number_of_L = 0
        number_of_R = 0
        balanced_splits = 0

        for c in s:
            if c == 'L':
                number_of_L += 1
            elif c == 'R':
                number_of_R += 1
            if number_of_L == number_of_R:
                balanced_splits += 1

        return balanced_splits


s = Solution()
print("Solution 1 :", s.balancedStringSplit("RLRRLLRLRL"))
print("Solution 2 :", s.balancedStringSplit("RLLLLRRRLR"))
print("Solution 3 :", s.balancedStringSplit("LLLLRRRR"))
print("Solution 4 :", s.balancedStringSplit("RLRRRLLRLL"))
