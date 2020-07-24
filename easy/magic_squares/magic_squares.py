class Solution(object):
    def numMagicSquaresInside(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        def is_magic_square(grid, i, j):
            elements = [grid[i][j], grid[i][j+1], grid[i][j+2], grid[i+1][j], grid[i+1][j+1],
                        grid[i+1][j+2], grid[i+2][j], grid[i+2][j+1], grid[i+2][j+2]]

            sorted_elements = [1, 2, 3, 4, 5, 6, 7, 8, 9]
            if not sorted(elements) == sorted_elements:
                return False

            first_row = grid[i][j] + grid[i][j+1] + grid[i][j+2]
            second_row = grid[i+1][j] + grid[i+1][j+1] + grid[i+1][j+2]
            third_row = grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]

            first_col = grid[i][j] + grid[i+1][j] + grid[i+2][j]
            second_col = grid[i][j+1] + grid[i+1][j+1] + grid[i+2][j+1]
            third_col = grid[i][j+2] + grid[i+1][j+2] + grid[i+2][j+2]

            first_diagonal = grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
            second_diagonal = grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j]

            if first_row == second_row == third_row == first_col == second_col == third_col == first_diagonal == second_diagonal:
                return True

            return False

        magic_squares = 0
        for i in range(0, len(grid)-2):
            for j in range(0, len(grid[0])-2):
                if is_magic_square(grid, i, j):
                    magic_squares += 1
        return magic_squares


s = Solution()
print("Solution 1 : ", s.numMagicSquaresInside([[4, 3, 8, 4],
                                                [9, 5, 1, 9],
                                                [2, 7, 6, 2]]))
