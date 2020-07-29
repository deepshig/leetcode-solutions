class Solution(object):
    def findRestaurant(self, list1, list2):
        """
        :type list1: List[str]
        :type list2: List[str]
        :rtype: List[str]
        """
        restaurant_to_index1 = {}
        for index in range(0, len(list1)):
            restaurant_to_index1[list1[index]] = index

        min_index_sum = 3000
        fav_restaurants = []
        for index2 in range(0, len(list2)):
            restaurant = list2[index2]
            if restaurant in restaurant_to_index1:
                index_sum = restaurant_to_index1[restaurant] + index2
                if index_sum == min_index_sum:
                    fav_restaurants.append(restaurant)
                if index_sum < min_index_sum:
                    min_index_sum = index_sum
                    fav_restaurants = [restaurant]

        return fav_restaurants


s = Solution()

list1 = ["Shogun", "Tapioca Express", "Burger King", "KFC"]
list2 = ["Piatti", "The Grill at Torrey Pines",
         "Hungry Hunter Steakhouse", "Shogun"]
print("Solution 1 : ", s.findRestaurant(list1, list2))

list1 = ["Shogun", "Tapioca Express", "Burger King", "KFC"]
list2 = ["KFC", "Shogun", "Burger King"]
print("Solution 2 :", s.findRestaurant(list1, list2))
