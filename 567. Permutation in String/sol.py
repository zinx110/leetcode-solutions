class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        chars = {}
        for s in s1:
            chars[s] = chars.get(s, 0) + 1
        for l in range(len(s2) - len(s1) + 1):
            if s2[l] not in chars:
                continue
            r = l
            chars2 = {}
            while r < len(s2) and s2[r] in chars and r - l + 1 <= len(s1):
                chars2[s2[r]] = chars2.get(s2[r], 0) + 1
                if chars2[s2[r]] > chars[s2[r]]:
                    break
                r += 1
            if chars2 == chars:
                return True
        return False


# O(n * m) time complexity, O(m) space complexity, n = len(s2), m = len(s1)


class Solution2:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False

        s1Count, s2Count = [0] * 26, [0] * 26
        for i in range(len(s1)):
            s1Count[ord(s1[i]) - ord('a')] += 1
            s2Count[ord(s2[i]) - ord('a')] += 1

        matches = 0
        for i in range(26):
            matches += 1 if s1Count[i] == s2Count[i] else 0

        l = 0
        for r in range(len(s1), len(s2)):
            if matches == 26:
                return True

            index = ord(s2[r]) - ord('a')
            s2Count[index] += 1
            if s1Count[index] == s2Count[index]:
                matches += 1
            elif s1Count[index] + 1 == s2Count[index]:
                matches -= 1 


            index = ord(s2[l]) - ord('a')
            s2Count[index] -= 1
            if s1Count[index] == s2Count[index]:
                matches += 1
            elif s1Count[index] - 1 == s2Count[index]:
                matches -= 1
            l += 1
        
        return matches == 26