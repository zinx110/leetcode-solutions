from typing import Optional
class ListNode:
    def __init__(self, nxt: Optional[ListNode] = None):
       self.nxt = nxt

class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        s = set()
        for n in nums:
            s.add(n)

        starts = []
        for n in s:
            if n - 1 not in s:
                starts.append(n)
        res = 0
        for n in starts:
            curr = n
            currCount = 1
            while curr + 1 in s:
                curr = curr + 1
                currCount += 1
            if currCount > res:
                res = currCount


        return res