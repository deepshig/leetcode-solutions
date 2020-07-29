package main

import "fmt"

func getMapOfRestaurantToIndex(list []string) map[string]int {
	restaurantToIndex := make(map[string]int, len(list))
	for i, restaurant := range list {
		restaurantToIndex[restaurant] = i
	}
	return restaurantToIndex
}

func findRestaurant(list1 []string, list2 []string) []string {
	restaurantToIndex1 := getMapOfRestaurantToIndex(list1)

	minIndexSum := 3000
	favouriteRestaurants := make([]string, 0, len(list1))

	for index2, restaurant := range list2 {
		if index1, ok := restaurantToIndex1[restaurant]; ok {
			indexSum := index1 + index2
			if indexSum == minIndexSum {
				favouriteRestaurants = append(favouriteRestaurants, restaurant)
			} else if indexSum < minIndexSum {
				minIndexSum = indexSum
				favouriteRestaurants = []string{restaurant}
			}
		}
	}
	return favouriteRestaurants
}

func main() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println("Solution 1 :", findRestaurant(list1, list2))

	list1 = []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 = []string{"KFC", "Shogun", "Burger King"}
	fmt.Println("Solution 2 :", findRestaurant(list1, list2))
}
