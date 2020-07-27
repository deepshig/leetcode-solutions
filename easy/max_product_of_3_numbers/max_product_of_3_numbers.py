class Solution(object):
    def maximumProduct(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        first_min, second_min = float("inf"), float("inf")
        first_max, second_max, third_max = - \
            float("inf"), -float("inf"), -float("inf")

        for n in nums:
            if n <= first_min:
                second_min, first_min = first_min, n
            elif n <= second_min:
                second_min = n

            if n >= first_max:
                third_max, second_max, first_max = second_max, first_max, n
            elif n >= second_max:
                third_max, second_max = second_max, n
            elif n >= third_max:
                third_max = n

        product_1 = first_min*second_min*first_max
        product_2 = first_max*second_max*third_max
        return max(product_1, product_2)


s = Solution()
print("Solution 1 : ", s.maximumProduct([1, 2, 3]))
print("Solution 2 : ", s.maximumProduct([1, 2, 3, 4]))
print("Solution 3 : ", s.maximumProduct([-4, -3, -2, -1, 60]))
