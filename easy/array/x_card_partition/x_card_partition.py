class Solution(object):
    def __gcd(self, a, b):
        if a == 0:
            return b
        return self.__gcd(b % a, a)

    def hasGroupsSizeX(self, deck):
        """
        :type deck: List[int]
        :rtype: bool
        """
        occurrences = {}
        for c in deck:
            if not c in occurrences:
                occurrences[c] = deck.count(c)

        x = occurrences[deck[0]]
        for _, o in occurrences.items():
            if self.__gcd(x, o) < 2:
                return False
            x = self.__gcd(x, o)

        return True


s = Solution()
print("Solution 1 : ", s.hasGroupsSizeX([1, 2, 3, 4, 4, 3, 2, 1]))
print("Solution 2 : ", s.hasGroupsSizeX([1, 1, 1, 2, 2, 2, 3, 3]))
print("Solution 3 : ", s.hasGroupsSizeX([1]))
print("Solution 4 : ", s.hasGroupsSizeX([1, 1]))
print("Solution 5 : ", s.hasGroupsSizeX([1, 1, 2, 2, 2, 2]))
