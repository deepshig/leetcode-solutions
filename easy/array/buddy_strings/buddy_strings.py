class Solution(object):
    def buddyStrings(self, A, B):
        """
        :type A: str
        :type B: str
        :rtype: bool
        """
        if len(A) != len(B):
            return False

        places_of_difference = []
        for i in range(0, len(A)):
            if A[i] != B[i]:
                places_of_difference.append(i)
            if len(places_of_difference) > 2:
                return False

        if len(places_of_difference) == 2:
            first, second = places_of_difference[0], places_of_difference[1]
            if A[first] != B[second] or A[second] != B[first]:
                return False

        if len(places_of_difference) == 1:
            return False

        if len(places_of_difference) == 0:
            for c in A:
                if A.count(c) >= 2:
                    return True
            return False

        return True


s = Solution()
print("Solution 1 : ", s.buddyStrings("ab", "ba"))
print("Solution 2 : ", s.buddyStrings("ab", "ab"))
print("Solution 3 : ", s.buddyStrings("aa", "aa"))
print("Solution 4 : ", s.buddyStrings("aaaaaaabc", "aaaaaaacb"))
print("Solution 5 : ", s.buddyStrings("", "aa"))
