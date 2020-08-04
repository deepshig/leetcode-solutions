class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        s_to_t, t_to_s = {}, {}
        for i, s_char in enumerate(s):
            if not s_char in s_to_t and not t[i] in t_to_s:
                s_to_t[s_char] = t[i]
                t_to_s[t[i]] = s_char
            elif not s_char in s_to_t or not t[i] in t_to_s:
                return False
            elif t_to_s[t[i]] != s_char or s_to_t[s_char] != t[i]:
                return False

        return True


s = Solution()
print("Solution 1 : ", s.isIsomorphic("egg", "add"))
print("Solution 2 : ", s.isIsomorphic("foo", "bar"))
print("Solution 3 : ", s.isIsomorphic("paper", "title"))
