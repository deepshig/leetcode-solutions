class Solution(object):
    def minDeletionSize(self, A):
        """
        :type A: List[str]
        :rtype: int
        """
        columns_to_delete = 0
        for i in range(0, len(A[0])):
            for j in range(1, len(A)):
                if A[j][i] < A[j-1][i]:
                    columns_to_delete += 1
                    break
        return columns_to_delete


s = Solution()
print("Solution 1 :", s.minDeletionSize(["cba", "daf", "ghi"]))
print("Solution 2 :", s.minDeletionSize(["a", "b"]))
print("Solution 3 :", s.minDeletionSize(["zyx", "wvu", "tsr"]))
print("Solution 4 :", s.minDeletionSize(["rrjk", "furt", "guzm"]))
