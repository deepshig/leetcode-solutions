class Solution(object):
    def rotate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        def reverse(arr, start, end):
            while start < end:
                arr[start], arr[end] = arr[end], arr[start]
                start, end = start+1, end-1

        elements = len(nums)
        k %= elements
        reverse(nums, 0, elements-1)
        reverse(nums, 0, k-1)
        reverse(nums, k, elements-1)
        return


s = Solution()

nums = [1, 2, 3, 4, 5, 6, 7]
s.rotate(nums, 3)
print("Solution 1 : ", nums)

nums = [-1, -100, 3, 99]
s.rotate(nums, 2)
print("Solution 2 : ", nums)

nums = [1, 2]
s.rotate(nums, 3)
print("Solution 3 : ", nums)

nums = [1, 2, 3]
s.rotate(nums, 2)
print("Solution 4 : ", nums)
