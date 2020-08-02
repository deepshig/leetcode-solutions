class Solution(object):
    def numJewelsInStones(self, J, S):
        """
        :type J: str
        :type S: str
        :rtype: int
        """
        jewels = 0
        for j in J:
            jewels += S.count(j)
        return jewels


s = Solution()
print("Solution 1 : ", s.numJewelsInStones("aA", "aAAbbbb"))
print("Solution 2 : ", s.numJewelsInStones("z", "ZZ"))
