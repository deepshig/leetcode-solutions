class Solution(object):
    def minTimeToVisitAllPoints(self, points):
        """
        :type points: List[List[int]]
        :rtype: int
        """
        current_pos, time = points[0], 0
        for p in points:
            if current_pos[0] != p[0] or current_pos[1] != p[1]:
                delta_x = abs(p[0]-current_pos[0])
                delta_y = abs(p[1]-current_pos[1])

                time += max(delta_x, delta_y)
            current_pos = p

        return time


s = Solution()
print("Solution 1 : ", s.minTimeToVisitAllPoints([[1, 1], [3, 4], [-1, 0]]))
print("Solution 2 : ", s.minTimeToVisitAllPoints([[3, 2], [-2, 2]]))
