class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:

        longest = 0
        seen = set()
        start = 0
        for c in s:
            while c in seen:
                seen.remove(s[start])
                start += 1
            seen.add(c)
            if len(seen) > longest:
                longest = len(seen)
        return longest

                
