from typing import List


class Solution:
    def powerfulIntegers(self, x: int, y: int, bound: int) -> List[int]:
        powerfulIntsSet = set()
        i = 1
        while i < bound:
            j = 1
            while i+j <= bound:
                powerfulIntsSet.add(i+j)
                if y == 1:
                    break
                j *= y

            if x == 1:
                break
            i *= x

        return list(powerfulIntsSet)


s = Solution()
print("Solution 1 : ", s.powerfulIntegers(2, 3, 10))
print("Solution 2 : ", s.powerfulIntegers(3, 5, 15))
