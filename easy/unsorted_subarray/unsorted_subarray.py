class Solution(object):
    def findUnsortedSubarray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        start, end = 0, len(nums)-1
        need_to_sort = False

        for i in range(0, len(nums)-1):
            if nums[i] > nums[i+1]:
                start = i
                need_to_sort = True
                break

        for i in range(len(nums)-1, 0, -1):
            if nums[i] < nums[i-1]:
                end = i
                need_to_sort = True
                break

        unsorted_subarray = nums[start:end+1]
        min_val, max_val = min(unsorted_subarray), max(unsorted_subarray)

        if not need_to_sort:
            return 0

        while start > 0 and nums[start-1] > min_val:
            start -= 1

        while end < len(nums)-1 and nums[end+1] < max_val:
            end += 1

        return end-start+1


s = Solution()
print("Solution 1 : ", s.findUnsortedSubarray([2, 6, 4, 8, 10, 9, 15]))
print("Solution 2 : ", s.findUnsortedSubarray([1, 2, 3, 4]))
print("Solution 3 : ", s.findUnsortedSubarray([1, 3, 2, 2, 2]))
