class Solution(object):
    def destCity(self, paths):
        """
        :type paths: List[List[str]]
        :rtype: str
        """
        cities_with_outgoing_paths = {}
        for path in paths:
            cities_with_outgoing_paths[path[0]] = True
            if not path[1] in cities_with_outgoing_paths:
                cities_with_outgoing_paths[path[1]] = False

        for city, is_source in cities_with_outgoing_paths.items():
            if not is_source:
                return city


s = Solution()
print("Solution 1 : ", s.destCity(
    [["London", "New York"], ["New York", "Lima"], ["Lima", "Sao Paulo"]]))
print("Solution 2 : ", s.destCity([["B", "C"], ["D", "B"], ["C", "A"]]))
print("Solution 3 : ", s.destCity([["A", "Z"]]))
