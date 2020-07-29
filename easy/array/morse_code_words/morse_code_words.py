class Solution(object):
    def uniqueMorseRepresentations(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        morse_codes = [".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..",
                       "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."]

        transformations = {}
        for word in words:
            morse_code = ""
            for letter in word:
                morse_code += morse_codes[ord(letter) - ord('a')]
            transformations[morse_code] = True

        return len(transformations)


s = Solution()
print("Solution 1 : ", s.uniqueMorseRepresentations(
    ["gin", "zen", "gig", "msg"]))
