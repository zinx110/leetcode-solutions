class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        chars = {}
        res = 0
        l = 0
        maxf = 0
        for r in range(len(s)):
            chars[s[r]] = chars.get(s[r], 0) + 1
            maxf = max(maxf, chars[s[r]])
            # if (r - l + 1) - max(chars.values()) > k:
            if (r - l + 1) - maxf > k:
                chars[s[l]] -= 1
                l += 1
            res = max(r - l + 1, res)
        return res



