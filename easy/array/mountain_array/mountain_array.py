class Solution(object):
    def validMountainArray(self, A):
        """
        :type A: List[int]
        :rtype: bool
        """
        if len(A) < 3 or A[0] >= A[1]:
            return False

        peak_acheived = False
        for i in range(1, len(A)):
            if A[i] == A[i-1]:
                return False
            if not peak_acheived:
                if A[i] < A[i-1]:
                    peak_acheived = True
            else:
                if A[i] > A[i-1]:
                    return False

        if not peak_acheived:
            return False
        return True


s = Solution()
print("Solution 1 : ", s.validMountainArray([2, 1]))
print("Solution 2 : ", s.validMountainArray([3, 5, 5]))
print("Solution 3 : ", s.validMountainArray([0, 3, 2, 1]))
print("Solution 4 : ", s.validMountainArray([1, 3, 2]))
