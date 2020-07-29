class Solution(object):
    def busyStudent(self, startTime, endTime, queryTime):
        """
        :type startTime: List[int]
        :type endTime: List[int]
        :type queryTime: int
        :rtype: int
        """
        students_studying = 0
        for i in range(0, len(endTime)):
            if endTime[i] >= queryTime and startTime[i] <= queryTime:
                students_studying += 1
        return students_studying


s = Solution()
print("Solution 1 :", s.busyStudent([1, 2, 3], [3, 2, 7], 4))
print("Solution 2 :", s.busyStudent([4], [4], 4))
print("Solution 3 :", s.busyStudent([4], [4], 5))
print("Solution 4 :", s.busyStudent([1, 1, 1, 1], [1, 3, 2, 4], 7))
print("Solution 5 :", s.busyStudent(
    [9, 8, 7, 6, 5, 4, 3, 2, 1], [10, 10, 10, 10, 10, 10, 10, 10, 10], 5))
