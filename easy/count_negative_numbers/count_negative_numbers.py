class Solution(object):
    def countNegatives(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        columns = len(grid[0])
        total_negatives = 0

        for row in grid:
            if (row[0] < 0):
                total_negatives += columns
                continue
            if (row[columns-1] >= 0):
                continue
            for val in reversed(row):
                if (val < 0):
                    total_negatives += 1
                else:
                    break

        return total_negatives


s = Solution()
print("Solution 1 :", s.countNegatives(
    [[4, 3, 2, -1], [3, 2, 1, -1], [1, 1, -1, -2], [-1, -1, -2, -3]]))
print("Solution 2 :", s.countNegatives([[3, 2], [1, 0]]))
print("Solution 3 :", s.countNegatives([[1, -1], [-1, -1]]))
print("Solution 4 :", s.countNegatives([[-1]]))
