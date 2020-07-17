class Solution(object):
    def checkPossibility(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        place_of_violation = -1

        for i in range(0, len(nums)-1):
            if nums[i] > nums[i+1]:
                if place_of_violation != -1:
                    return False
                place_of_violation = i

        if place_of_violation == -1:
            return True

        index_to_change = place_of_violation
        if index_to_change == 0:
            return True
        elif nums[index_to_change-1] <= nums[index_to_change+1]:
            return True

        index_to_change = place_of_violation + 1
        if index_to_change == len(nums)-1:
            return True
        elif nums[index_to_change-1] <= nums[index_to_change+1]:
            return True

        return False
