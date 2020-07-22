class Solution(object):
    def __flood_fill_recursion(self, image, sr, sc, old_color, newColor):
        image[sr][sc] = newColor

        if sr > 0 and image[sr-1][sc] == old_color:
            image = self.__flood_fill_recursion(
                image, sr-1, sc, old_color, newColor)

        if sr < len(image)-1 and image[sr+1][sc] == old_color:
            image = self.__flood_fill_recursion(
                image, sr+1, sc, old_color, newColor)

        if sc > 0 and image[sr][sc-1] == old_color:
            image = self.__flood_fill_recursion(
                image, sr, sc-1, old_color, newColor)

        if sc < len(image[0])-1 and image[sr][sc+1] == old_color:
            image = self.__flood_fill_recursion(
                image, sr, sc+1, old_color, newColor)
        return image

    def floodFill(self, image, sr, sc, newColor):
        """
        :type image: List[List[int]]
        :type sr: int
        :type sc: int
        :type newColor: int
        :rtype: List[List[int]]
        """
        old_color = image[sr][sc]
        if old_color == newColor:
            return image
        return self.__flood_fill_recursion(image, sr, sc, old_color, newColor)


s = Solution()
print("Solution 1 : ", s.floodFill([[1, 1, 1], [1, 1, 0], [1, 0, 1]], 1, 1, 2))
print("Solution 1 : ", s.floodFill([[0, 0, 0], [0, 1, 1]], 1, 1, 1))
