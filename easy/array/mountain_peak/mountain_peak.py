from typing import List


class Solution:
    def peakIndexInMountainArray(self, A: List[int]) -> int:
        start, end = 0, len(A)
        while start < end:
            mid = int((start+end)/2)
            if A[mid-1] < A[mid] and A[mid] > A[mid+1]:
                return mid
            elif A[mid+1] > A[mid]:
                start = mid+1
            elif A[mid-1] > A[mid]:
                end = mid


s = Solution()
print("Solution 1 : ", s.peakIndexInMountainArray([0, 1, 0]))
print("Solution 2 : ", s.peakIndexInMountainArray([0, 2, 1, 0]))
