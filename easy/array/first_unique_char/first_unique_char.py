class Solution:
    def firstUniqChar(self, s: str) -> int:
        char_to_index = [0]*26
        for c in s:
            char_to_index[ord(c)-ord('a')] += 1

        for i, c in enumerate(s):
            if char_to_index[ord(c)-ord('a')] == 1:
                return i

        return -1


s = Solution()
print("Solution 1 : ", s.firstUniqChar("leetcode"))
print("Solution 2 : ", s.firstUniqChar("loveleetcode"))
