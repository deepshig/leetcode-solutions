class Solution:
    def wordPattern(self, pattern: str, str: str) -> bool:
        char_to_word, word_to_char = {}, {}
        words = str.split(" ")
        if len(pattern) != len(words):
            return False

        for i in range(0, len(words)):
            if not pattern[i] in char_to_word and not words[i] in word_to_char:
                char_to_word[pattern[i]] = words[i]
                word_to_char[words[i]] = pattern[i]
            elif not pattern[i] in char_to_word or not words[i] in word_to_char:
                return False
            elif char_to_word[pattern[i]] != words[i] or word_to_char[words[i]] != pattern[i]:
                return False
        return True


s = Solution()
print("Solution 1 : ", s.wordPattern("abba", "dog cat cat dog"))
print("Solution 2 : ", s.wordPattern("abba", "dog cat cat fish"))
print("Solution 3 : ", s.wordPattern("aaaa", "dog cat cat fish"))
print("Solution 4 : ", s.wordPattern("abba", "dog dog dog fish"))
